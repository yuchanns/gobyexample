package main

import (
	"flag"
	"github.com/yuchanns/gobyexample/blog/server"
	"github.com/yuchanns/gobyexample/utils/uviper"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	flag.Parse()

	uviper.Init()

	closeFunc, err := server.AppInit()

	if err != nil {
		panic(err.Error())
	}

	log.Print("blog start")

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	for {
		s := <-c
		switch s {
		case syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT:
			closeFunc()
			log.Println("blog stop")
			return
		default:
			return
		}
	}
}
