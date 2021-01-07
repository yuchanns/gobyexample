package server

import (
	"context"
	"github.com/yuchanns/gobyexample/grpc-app/app"
	hello_world "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
	"github.com/yuchanns/gobyexample/grpc-app/server/pb"
)

type GreeterServer struct {
	svc *app.GreeterSvc
}

func NewGreeterServer(svc *app.GreeterSvc) hello_world.GreeterServer {
	return &GreeterServer{svc: svc}
}

func (g *GreeterServer) SayHello(
	ctx context.Context,
	request *hello_world.HelloRequest,
) (
	*hello_world.HelloResponse,
	error,
) {
	dto := pb.PbToHelloDTO(request)
	message := g.svc.SayHello(ctx, dto)
	return &hello_world.HelloResponse{
		Message: message,
	}, nil
}
