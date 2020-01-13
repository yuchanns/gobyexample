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

	markdown := engine.Group("/markdown")
	{
		markdown.POST("/add", srv.Add)
		markdown.GET("/get", srv.Get)
	}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}
