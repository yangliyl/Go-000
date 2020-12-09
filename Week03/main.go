package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

func main() {
	serve := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		select {
		case <-stop:
			log.Println("signal")
		case <-ctx.Done():
			log.Println("context done")
		}
		log.Println("shutting down")
		return serve.Shutdown(ctx)
	})

	g.Go(func() error {
		log.Println("server start")
		return serve.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Printf("%+v\n", err)
	}
}
