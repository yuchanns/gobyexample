package main

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", pongo2.Context{
		"greet": "hello",
		"obj":   "world",
	})
}
