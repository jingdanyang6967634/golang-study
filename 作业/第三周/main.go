package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/sync/errgroup"
)

func StartHttpServer(srv *http.Server) error {
	// Hello world, the web server
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)

	return srv.ListenAndServe()
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, errCtx := errgroup.WithContext(ctx)

	var srv *http.Server = &http.Server{Addr: "127.0.0.1:8080"}

	g.Go(func() error {
		return StartHttpServer(srv)
	})

	g.Go(func() error {
		<-errCtx.Done()
		return srv.Shutdown(errCtx)
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c)

	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-c:
				cancel()
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")
}
