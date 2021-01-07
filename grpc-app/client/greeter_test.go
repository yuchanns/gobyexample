package client

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/yuchanns/gobyexample/grpc-app/common"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestSayHello(t *testing.T) {
	tracer, closer, err := common.NewJaegerTracer("grpc-app", os.Getenv("AGENT_HOST_PORT"))
	if err != nil {
		t.Fatalf("failed to create jaeger tracer")
	}
	defer closer.Close()
	conn, err := grpc.Dial(":9090",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(tracer, otgrpc.LogPayloads())),
		grpc.WithStreamInterceptor(otgrpc.OpenTracingStreamClientInterceptor(tracer, otgrpc.LogPayloads())),
	)
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
