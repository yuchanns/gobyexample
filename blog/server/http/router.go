package http

import "github.com/gin-gonic/gin"

func SetSrvContext(c *gin.Context) {
	srv.SetContext(c)

	c.Next()
}

func InitRouter(engine *gin.Engine) {
	engine.Use(SetSrvContext)

	p := engine.Group("/ping")
	{
		p.GET("/", Ping)
	}
}

func Ping(*gin.Context) {
	srv.Ping()
}
