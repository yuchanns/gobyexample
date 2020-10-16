package main

import (
	"context"
	"fmt"
	proto "github.com/yuchanns/gobyexample/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
)

const port = ":52242"

type Server struct{}

func (s *Server) SayHello(_ context.Context, in *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Printf("get request from client: %v\n", in.Name)
	return &proto.HelloResponse{
		Message: strings.Join([]string{
			"Hello",
			in.Name,
		}, " "),
	}, nil
}

func NewServer() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	srv := grpc.NewServer()
	proto.RegisterGreeterServer(srv, &Server{})
	reflection.Register(srv)
	if err := srv.Serve(l); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	NewServer()
}
