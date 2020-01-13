package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/blog/service"
	"log"
	"net/http"
	"time"
)

var (
	srv *service.Service
)

func New(s *service.Service) func() {
	srv = s

	engine := gin.Default()

	InitRouter(engine)

	server := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatal("Http Start:", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	cancelFunc := func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Fatal("Http Shutdown:", err)
		}

		cancel()

		log.Fatal("Http Exiting")
	}

	return cancelFunc
}
