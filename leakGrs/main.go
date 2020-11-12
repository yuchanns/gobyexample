package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func leakGrs() error {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	ch := make(chan error)

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
	engine.GET("/", handlerLeakGrs)
	http.Handle("/", engine)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("cannot start server: %+v", err)
	}
}
