package main

import (
	"context"
	"log"
	"moodbase-server/handlers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "moodbase-server", log.LstdFlags)

	healthHandler := handlers.NewHealth(logger)
	productHandler := handlers.NewProductHandler(logger)

	sm := http.NewServeMux()
	sm.Handle("/health", healthHandler)
	sm.Handle("/products", productHandler)

	s := &http.Server{
		Addr:         "127.0.0.1:9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		logger.Println("Starting server...")
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal()
		}
	}()

	// Signal channel
	// Kill the server when the channel receives appropriate signal e.g. Ctrl C in cmd
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, graceful shutdown", sig)

	// Graceful shutdown. Useful for cases like server upgrades.
	// Waits for existing requests to finish before shutdown.
	timeoutContext, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(timeoutContext)


}