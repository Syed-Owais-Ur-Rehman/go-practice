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

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Practicing Graceful Shutdown with Mux server.")

	// Set up the Gorilla Mux router
	r := mux.NewRouter()

	// Define a simple handler
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// Create an HTTP server with the router
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Set up signal channel to listen for interrupt or termination signals
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Run the server in a goroutine
	go func() {
		log.Println("Server is starting...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error starting server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down
	sig := <-sigs
	log.Printf("Received signal: %v. Shutting down gracefully...\n", sig)

	// Set a timeout for the graceful shutdown (e.g., 5 seconds)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // cancel variable holds the memory address of the cancel function.
	defer cancel()

	// Call the graceful shutdown method
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server Shutdown failed: %v", err)
	}
	log.Println("Server exited cleanly")
}
