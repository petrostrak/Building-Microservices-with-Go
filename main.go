package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/petrostrak/Building-Microservices-with-Go/product-api/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProduct(l)

	mux := http.NewServeMux()
	mux.Handle("/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 2)

	// signal.Notify will broadcast a message on this channel whenever
	// an os.Kill / os.Interrupt signal is received.
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
