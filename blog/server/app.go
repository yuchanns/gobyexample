package server

import (
	"github.com/yuchanns/gobyexample/blog/dao"
	"github.com/yuchanns/gobyexample/blog/server/http"
	"github.com/yuchanns/gobyexample/blog/service"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	d := dao.New()
	srv := service.New(d)
	cancelFunc := http.New(srv)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)

	for {
		s := <-quit
		switch s {
		case syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM:
			cancelFunc()
			return
		default:
			return
		}
	}
}
