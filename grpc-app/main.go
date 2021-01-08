package main

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/infra/startup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	tracer, closer, err := common.NewJaegerTracer("grpc-app", os.Getenv("AGENT_HOST_PORT"))
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

	if err := startup.RegisterGrpcServer(srv); err != nil {
		log.Fatalf("failed to register grpc server: %+v", err)
	}

	if err := startup.RegisterVars(); err != nil {
		log.Fatalf("failed to register vars: %+v", err)
	}

	reflection.Register(srv)

	log.Println("start at :8080")

	if err := srv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
