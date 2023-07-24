package server

import (
	"context"
	"go-app/config"
	"go-app/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// Run starts the HTTP server.
func Run() {
	// Get the server configuration options from environment variables.
	serverConfigs := config.GetServerConfigs()

	// Initialize the router with the application's routes.
	router := router.SetupRouter(serverConfigs.ServiceName)

	// Create an HTTP server with the specified address and router.
	server := &http.Server{
		Addr:    serverConfigs.Port,
		Handler: router,
	}

	// Start the server with or without graceful shutdown.
	if serverConfigs.GracefulShutdown {
		startWithGracefulShutdown(serverConfigs.ServiceName, server)
	} else {
		start(serverConfigs.ServiceName, server)
	}
}

// start starts the HTTP server without graceful shutdown.
func start(serviceName string, server *http.Server) {
	log.Printf("Started %s service on port %s", serviceName, server.Addr)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}
}

// startWithGracefulShutdown starts the HTTP server with graceful shutdown.
func startWithGracefulShutdown(serviceName string, server *http.Server) {
	log.Printf("Started %s service on port %s with graceful shutdown ", serviceName, server.Addr)

	// Create a channel to signal when all idle connections are closed.
	idleConnsClosed := make(chan struct{})

	// Start a goroutine to listen for interrupts and shut down the server gracefully when one is received.
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		log.Println("Shutting down server...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("Error shutting down server: %s\n", err)
		}

		close(idleConnsClosed)
	}()

	// Start the server and wait for it to return.
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Error starting server: %s\n", err)
	}

	<-idleConnsClosed
	log.Println("Server stopped")
}
