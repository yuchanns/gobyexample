package main

import (
	"github.com/gomodule/redigo/redis"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	queue := &Queue{conn: conn}

	msg := &Message{
		name: "demoQueue",
	}
	queue.InitReceiver(msg, 1)

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT)

	for {
		switch <-quit {
		case syscall.SIGINT:
			os.Exit(0)
		default:
			return
		}
	}
}
