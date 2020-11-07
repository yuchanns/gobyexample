package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	hello := engine.Group("/hello")
	hello.GET("/", func(ctx *gin.Context) {
		go func() {
			for {
				//
			}
		}()
		ctx.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/", engine)
	if err := http.ListenAndServe(":9001", nil); err != nil {
		log.Fatalf("failed to start server: %s", err)
	}
}
