package main

import (
	"flag"
	"github.com/yuchanns/gobyexample/blog/server"
	"github.com/yuchanns/gobyexample/utils/uviper"
)

func main() {
	flag.Parse()

	uviper.Init()

	server.Init()
}
