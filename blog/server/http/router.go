package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(engine *gin.Engine) {
	test := engine.Group("/test")
	{
		test.GET("/ping", Ping)
	}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}
