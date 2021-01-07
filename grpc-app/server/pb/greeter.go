package pb

import (
	dto2 "github.com/yuchanns/gobyexample/grpc-app/app/dto"
	"github.com/yuchanns/gobyexample/grpc-app/common"
	helloworld "github.com/yuchanns/gobyexample/grpc-app/proto/greeter"
)

func PbToHelloDTO(pb *helloworld.HelloRequest) *dto2.HelloDTO {
	dto := &dto2.HelloDTO{}
	common.MustConvert(pb, &dto)
	return dto
}
