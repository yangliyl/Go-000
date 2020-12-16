package main

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"week04/api/demo"

	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	// init config

	// init db
	db := &sql.DB{}

	server := grpc.NewServer()

	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		stop := make(chan os.Signal)
		signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		select {
		case <-stop:
			log.Println("signal")
		case <-ctx.Done():
			log.Println(ctx.Err())
		}
		log.Println("shutting down")
		server.GracefulStop()
		return errors.New("signal")
	})

	g.Go(func() error {
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			os.Exit(1)
		}

		service := InitializeDemo(db)

		demo.RegisterUserServer(server, service)
		return server.Serve(listener)
	})

	if err := g.Wait(); err != nil {
		log.Println("%+v\n", err)
	}
}
