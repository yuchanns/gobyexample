package main

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", Index)
	r.HTMLRender = New("pongo2render/templates")

	return r
}
