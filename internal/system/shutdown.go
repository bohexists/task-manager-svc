package shutdown

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

// ShutdownConfig structure for shutdown
type ShutdownConfig struct {
	GrpcServer      *grpc.Server
	NATSConnection  *nats.Conn
	ShutdownTimeout time.Duration
}

// ListenForShutdown gracefully shuts down the system
func ListenForShutdown(cfg ShutdownConfig) {
	// Channel for graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Wait for shutdown signal
	<-quit
	log.Println("Shutdown signal received, gracefully shutting down...")

	// Create context for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
	defer cancel()

	// Close gRPC server
	if cfg.GrpcServer != nil {
		cfg.GrpcServer.GracefulStop()
		log.Println("gRPC server shut down successfully")
	}

	// Close NATS connection
	if cfg.NATSConnection != nil {
		cfg.NATSConnection.Close()
		log.Println("NATS connection closed successfully")
	}

	// wait for shutdown to complete
	<-ctx.Done()
	if ctx.Err() == context.DeadlineExceeded {
		log.Println("Shutdown timed out, forcing exit.")
	} else {
		log.Println("Graceful shutdown completed successfully.")
	}
}
