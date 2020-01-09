package http

import "github.com/gin-gonic/gin"

func InitRouter(engine *gin.Engine) {
	p := engine.Group("/ping")
	{
		p.GET("/", Ping)
	}

	m := engine.Group("/markdown")
	{
		m.POST("/add", Add)
	}
}

func Ping(c *gin.Context) {
	srv.Ping(c)
}

func Add(c *gin.Context) {
	srv.AddMarkdown(c)
}
