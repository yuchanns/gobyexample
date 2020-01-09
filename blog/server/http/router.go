package http

import "github.com/gin-gonic/gin"

func InitRouter(engine *gin.Engine) {
	p := engine.Group("/ping")
	{
		p.GET("/", srv.Ping)
	}

	m := engine.Group("/markdown")
	{
		m.POST("/add", srv.AddMarkdown)
		m.GET("/get", srv.GetMarkdown)
	}
}
