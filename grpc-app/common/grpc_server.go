package common

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
)

type GrpcApplication struct {
	Endpoint           string
	GatewayAddr        string
	AppName            string
	AgentHostPort      string
	RegisterGrpcServer func(srv *grpc.Server) error
	RegisterGateway    func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error
	RegisterVars       func() error
}

type GrpcServer struct {
	app *GrpcApplication
	mux *runtime.ServeMux
	l   net.Listener
	srv *grpc.Server
}

func NewGrpcServer(app *GrpcApplication) *GrpcServer {
	l, err := net.Listen("tcp", app.Endpoint)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer, closer, err := NewJaegerTracer(app.AppName, app.AgentHostPort)
	if err == nil {
		defer closer.Close()
	} else {
		log.Println("failed to create Jaeger tracer")
	}

	opts := []grpc.ServerOption{
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
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
		mux: mux,
		l:   l,
		srv: srv,
		app: app,
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
