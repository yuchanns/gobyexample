package json_iterator

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"net/http"
)

func SetupRoter() *gin.Engine {
	// to build with jsoniter, a build pragma should be in the main.go file
	// such as "// +build jsoniter"
	jsoniter.RegisterTypeEncoder("json_iterator.Location", &locationAsStringCodec{})
	r := gin.Default()
	r.GET("/jsoniter", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": &s,
		})
	})

	return r
}
