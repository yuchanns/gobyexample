package main

import (
	"errors"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func leakGrs() error {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan error, 9)

	for i := range s {
		go func(i int) {
			var err error
			if i == 3 {
				err = errors.New("something wrong")
			}
			ch <- err
		}(i)
	}

	for range s {
		err := <-ch
		if err != nil {
			return err
		}
	}

	return nil
}

func handlerLeakGrs(c *gin.Context) {
	err := leakGrs()
	c.JSON(http.StatusOK, gin.H{
		"err": fmt.Sprintf("%+v", err),
	})
}

func main() {
	engine := gin.Default()
	pprof.Register(engine)
	engine.GET("/", handlerLeakGrs)
	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %+v", err)
	}
}
