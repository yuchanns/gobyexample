package app

import (
	"context"
	dto2 "github.com/yuchanns/gobyexample/grpc-app/app/dto"
	"github.com/yuchanns/gobyexample/grpc-app/domain/greeter"
)

type GreeterSvc struct {
	greeterDS greeter.IDomSvc
}

func NewGreeterSvc(greeterDS greeter.IDomSvc) *GreeterSvc {
	return &GreeterSvc{greeterDS: greeterDS}
}

func (s *GreeterSvc) SayHello(ctx context.Context, dto *dto2.HelloDTO) string {
	do := dto2.DTOtoGreeter(dto)
	return s.greeterDS.SayHello(ctx, do)
}
