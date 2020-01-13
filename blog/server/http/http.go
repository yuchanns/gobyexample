package http

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/blog/server/uviper"
	"github.com/yuchanns/gobyexample/blog/service"
	"log"
	"net/http"
	"time"
)

var (
	srv *service.Service
)

type httpConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func New(s *service.Service) func() {
	var config httpConfig
	if err := uviper.Get("http").Unmarshal(&config); err != nil {
		panic(fmt.Sprintln("Read httpconfig failed:", err))
	}

	srv = s

	engine := gin.Default()

	InitRouter(engine)

	server := &http.Server{
		Addr:         config.Addr,
		Handler:      engine,
		ReadTimeout:  config.ReadTimeout * time.Second,
		WriteTimeout: config.WriteTimeout * time.Second,
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
