package _interface

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/utils/stdresp"
)

type IService interface {
	Json(*gin.Context, *stdresp.DefaultResp)
	Ping(*gin.Context)
}
