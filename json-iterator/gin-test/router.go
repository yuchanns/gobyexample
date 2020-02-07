package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/student", func(c *gin.Context) {
		s := Student{
			ID:     1,
			Age:    27,
			Gender: 1,
			Name:   "yuchanns",
			Location: Location{
				Country:  "China",
				Province: "Guangdong",
				City:     "Shenzhen",
				District: "Nanshan",
			},
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": &s,
		})
	})

	r.GET("/location", func(c *gin.Context) {
		l := Location{
			Country:  "China",
			Province: "Guangdong",
			City:     "Shenzhen",
			District: "Nanshan",
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": &l,
		})
	})

	return r
}
