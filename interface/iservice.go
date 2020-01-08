package _interface

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/utils/stdresp"
)

type IService interface {
	SetContext(*gin.Context)
	Json(*stdresp.DefaultResp)
	Ping()
}
