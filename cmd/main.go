package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln("main: error: ", err.Error())
	}
}

type dummy struct{}

func (d *dummy) ServeHTTP(w http.ResponseWriter, r *http.Request) {}

func run() error {
	log := log.New(os.Stdout, ">> ", log.LstdFlags|log.Lshortfile)
	log.Println("initializing service")

	server := http.Server{
		Addr:         ":8080",
		Handler:      handlers.NewAPI()
		ErrorLog:     log,
		ReadTimeout:  time.Duration(1) * time.Minute, // TODO: get from configuration
		WriteTimeout: time.Duration(1) * time.Minute, // TODO: get from configuration
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	serverErr := make(chan error, 1)
	go func() {
		log.Println("server listening on 8080")
		serverErr <- server.ListenAndServe()
	}()

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErr:
		return fmt.Errorf("server error %w", err)
	case sig := <-shutdown:
		log.Println("shutdown signal called")

		// Give outstanding requests ample time to finish
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(2)*time.Minute)
		defer cancel()

		err := server.Shutdown(ctx)
		if err != nil {
			server.Close()
			log.Printf("graceful shutdown failed after %v %v", time.Duration(2)*time.Minute, err)
		}

		switch {
		case err != nil:
			return fmt.Errorf("could not gracefully shutdown %w", err)
		case sig == syscall.SIGSTOP:
			return fmt.Errorf("intrigrity issue caused shutdown %w", err)
		}
	}
	return nil
}
