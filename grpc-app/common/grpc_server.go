package common

import (
	"context"
	"github.com/go-playground/validator/v10"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
	"net/http"
)

type GrpcApplication struct {
	Endpoint           string `validate:"required"`
	GatewayAddr        string `validate:"required"`
	AppName            string `validate:"required"`
	AgentHostPort      string
	RegisterGrpcServer func(srv *grpc.Server) error
	RegisterGateway    func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	RegisterVars       func() error
}

func (a *GrpcApplication) Validate() error {
	vd := validator.New()
	if errs := vd.Struct(a); errs != nil {
		return errs
	}
	return nil
}

type GrpcServer struct {
	app     *GrpcApplication
	mux     *runtime.ServeMux
	l       net.Listener
	srv     *grpc.Server
	closers []io.Closer
}

func NewGrpcServer(app *GrpcApplication) *GrpcServer {
	if err := app.Validate(); err != nil {
		log.Fatalf("validation failed: %s", err)
	}
	l, err := net.Listen("tcp", app.Endpoint)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var closers []io.Closer

	tracer, closer, err := NewJaegerTracer(app.AppName, app.AgentHostPort)
	if err == nil {
		closers = append(closers, closer)
	} else {
		log.Println("failed to create Jaeger tracer")
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			otgrpc.OpenTracingStreamServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			otgrpc.OpenTracingServerInterceptor(tracer, otgrpc.LogPayloads()),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	}

	srv := grpc.NewServer(opts...)
	mux := runtime.NewServeMux()

	if app.RegisterGrpcServer != nil {
		if err := app.RegisterGrpcServer(srv); err != nil {
			log.Fatalf("failed to register grpc server: %+v", err)
		}
	}

	if app.RegisterGateway != nil {
		if err := app.RegisterGateway(context.Background(), mux, app.Endpoint, []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())),
			grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads())),
		}); err != nil {
			log.Fatalf("failed to register grpc gateway: %+v", err)
		}
	}

	if app.RegisterVars != nil {
		if err := app.RegisterVars(); err != nil {
			log.Fatalf("failed to register vars: %s", err)
		}
	}

	reflection.Register(srv)

	return &GrpcServer{
		mux:     mux,
		l:       l,
		srv:     srv,
		app:     app,
		closers: closers,
	}
}

func (s *GrpcServer) Run() {
	log.Printf("grpc server start at %s, gateway start at %s\n", s.app.Endpoint, s.app.GatewayAddr)

	go func() {
		if err := http.ListenAndServe(s.app.GatewayAddr, s.mux); err != nil {
			log.Fatalf("failed to start grpc gateway: %+v", err)
		}
	}()

	if err := s.srv.Serve(s.l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcServer) Close() error {
	for i := range s.closers {
		if err := s.closers[i].Close(); err != nil {
			return err
		}
	}

	return nil
}
