package main

import (
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"net/http"
	"os"
	"time"
)

const layout = "2006-01-02 15:04:05"

func main() {
	http.HandleFunc("/delay", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Second)
		str := fmt.Sprintf("hello world in %s and pid is %d", time.Now().Format(layout), os.Getpid())
		_, _ = writer.Write([]byte(str))
		log.Println(str)
	})
	srv := endless.NewServer(":9001", nil)
	srv.BeforeBegin = func(add string) {
		log.Printf("pid is %d", os.Getpid())
	}

	_ = srv.ListenAndServe()
}
