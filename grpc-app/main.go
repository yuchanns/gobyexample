package main

import (
	"github.com/yuchanns/gobyexample/grpc-app/infra/startup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":9090")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	if err := startup.RegisterGrpcServer(srv); err != nil {
		log.Fatalf("failed to register grpc server: %+v", err)
	}

	if err := startup.RegisterVars(); err != nil {
		log.Fatalf("failed to register vars: %+v", err)
	}

	reflection.Register(srv)

	if err := srv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
