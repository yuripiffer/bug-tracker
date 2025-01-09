package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bugtracker-backend/internal/db"
	"bugtracker-backend/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	// Ensure cleanup happens
	defer db.Cleanup()

	r := mux.NewRouter()

	// Register routes
	handlers.RegisterRoutes(r)

	// Create server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Channel to listen for errors coming from the listener
	serverErrors := make(chan error, 1)

	// Start the server
	go func() {
		log.Printf("Server starting on port %s...\n", srv.Addr)
		serverErrors <- srv.ListenAndServe()
	}()

	// Channel to listen for an interrupt or terminate signal from the OS
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown
	select {
	case err := <-serverErrors:
		log.Fatalf("Error starting server: %v", err)

	case sig := <-shutdown:
		log.Printf("Received signal %v: initiating shutdown", sig)

		// Give outstanding requests a deadline for completion
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		// Asking listener to shut down and shed load
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("Graceful shutdown did not complete in %v: %v", 10*time.Second, err)
			if err := srv.Close(); err != nil {
				log.Fatalf("Could not stop server gracefully: %v", err)
			}
		} else {
			log.Println("Server shut down gracefully.")
		}
	}
}
