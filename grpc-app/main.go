package main

import (
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

	var opts []grpc.ServerOption

	if middlewares, closeFunc, err := common.BuildGrpcOpentracingMiddlewares(
		"grpc-app",
		os.Getenv("AGENT_HOST_PORT"),
	); err == nil {
		defer closeFunc()
		opts = append(opts, middlewares...)
	} else {
		log.Println(err)
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
