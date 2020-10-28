package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yuchanns/gobyexample/ddd/infra/startup"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	if err := startup.RegisterHttpRoute(engine); err != nil {
		log.Fatalf("failed to register http route: %+v", err)
	}
	server := &http.Server{
		Addr:    ":9001",
		Handler: engine,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("failed to start http server: %+v", err)
	}
}
