package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

const layout = "2006-01-02 15:04:05"

var reload = flag.Bool("reload", false, "reload from fd 3")

func main() {
	flag.Parse()

	appRun()
}

func appRun() {
	http.HandleFunc("/delay", func(writer http.ResponseWriter, request *http.Request) {
		time.Sleep(20 * time.Second)
		str := fmt.Sprintf("hello world in %s and pid is %d", time.Now().Format(layout), os.Getpid())
		_, _ = writer.Write([]byte(str))
		log.Println(str)
	})

	var (
		ls  net.Listener
		err error
	)

	srv := http.Server{}

	if *reload {
		log.Printf("reload pid is %d", os.Getpid())
		log.Printf("graceful reload in %s\n", time.Now().Format(layout))
		fd := os.NewFile(3, "")
		ls, err = net.FileListener(fd)
	} else {
		log.Printf("start pid is %d", os.Getpid())
		log.Printf("start in %s\n", time.Now().Format(layout))
		ls, err = net.Listen("tcp", ":9001")
	}

	if err != nil {
		log.Fatalf("listen failed: %s\n", err)
	}

	go func() {
		err := srv.Serve(ls)
		if err != nil {
			log.Printf("serve failed: %s\n", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)

	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Println("stop")
			_ = srv.Shutdown(ctx)
			return
		case syscall.SIGUSR2:
			if err := graceful(ls); err != nil {
				log.Fatalf("graceful reload failed: %s\n", err)
			}
			_ = srv.Shutdown(ctx)
			return
		}
	}
}

func graceful(ls net.Listener) error {

	tl, ok := ls.(*net.TCPListener)
	if !ok {
		return errors.New("cannot assert listener into *net.TCPListener")
	}

	fd, err := tl.File()

	if err != nil {
		return err
	}

	cmd := exec.Command(os.Args[0], "-reload")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.ExtraFiles = []*os.File{fd}

	return cmd.Start()
}
