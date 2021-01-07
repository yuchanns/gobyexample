package client

import (
	"context"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"google.golang.org/grpc"
	"testing"
)

func TestSayHello(t *testing.T) {
	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		t.Fatalf("failed to connect: %v", err)
	}
	c := helloworld.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &helloworld.HelloRequest{
		Name: "yuchanns",
	})

	if err != nil {
		t.Fatalf("failed to SayHello: %v", err)
	}

	t.Logf("get response from SayHello: %v\n", r.Message)
}
