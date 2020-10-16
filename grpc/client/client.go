package main

import (
	"context"
	"fmt"
	proto "github.com/yuchanns/gobyexample/grpc/proto"
	"google.golang.org/grpc"
	"log"
)

const (
	address = "localhost:52242"
	name    = "yuchanns"
)

func NewClient() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	c := proto.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &proto.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("failed to SayHello: %v", err)
	}

	fmt.Printf("get response from SayHello: %v\n", r.Message)
}

func main() {
	NewClient()
}
