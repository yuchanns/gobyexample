package http

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/blog/service"
	_interface "github.com/yuchanns/gobyexample/interface"
	"github.com/yuchanns/gobyexample/utils/uviper"
	"log"
	"net/http"
	"time"
)

var srv *service.Service

type httpConfig struct {
	Addr         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func New(service *service.Service) (closeFunc func(), err error) {
	var _ _interface.IService = service

	srv = service

	var config httpConfig

	if err = uviper.Get("http").Unmarshal(&config); err != nil {
		return
	}

	engine := gin.Default()

	s := &http.Server{
		Addr:         config.Addr,
		Handler:      engine,
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	}

	InitRouter(engine)

	go func() {
		if err := s.ListenAndServe(); err != nil {
			panic(err.Error())
		}
	}()

	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)

	closeFunc = func() {
		cancelFunc()

		if err := srv.DB.Close(); err != nil {
			log.Fatal("close db failed:", err.Error())
		}

		if err := s.Shutdown(ctx); err != nil {
			log.Fatal("shut down failed:", err.Error())
		}
	}

	return
}
