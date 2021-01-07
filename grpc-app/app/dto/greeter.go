package dto

import (
	"github.com/yuchanns/gobyexample/grpc-app/common"
	"github.com/yuchanns/gobyexample/grpc-app/domain/greeter"
)

type HelloDTO struct {
	Name string `json:"name"`
}

func DTOtoGreeter(dto *HelloDTO) *greeter.Greeter {
	do := &greeter.Greeter{}
	common.MustConvert(dto, do)
	return do
}
