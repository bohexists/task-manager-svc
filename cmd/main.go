package main

import (
	"github.com/bohexists/task-manager-svc/internal/adapters/db"
	"github.com/bohexists/task-manager-svc/internal/adapters/grpc"
	"github.com/bohexists/task-manager-svc/internal/adapters/nats"
	"github.com/bohexists/task-manager-svc/internal/app"
	"github.com/bohexists/task-manager-svc/internal/config"
	shutdown "github.com/bohexists/task-manager-svc/internal/system"
	"github.com/bohexists/task-manager-svc/ports/inbound"
	"log"
	"time"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db.ConnectToDB(cfg)

	// Initialize repository and service
	taskRepo := db.NewTaskRepository(db.DB)
	taskService := app.NewTaskService(taskRepo)

	// Start gRPC server and get its instance for shutdown
	grpcServerInstance := grpc.StartGRPCServer(inbound.NewTaskServiceServer(taskService))

	// Start NATS subscriber and get its connection for shutdown
	natsConn, err := nats.InitNATSSubscriber(cfg, taskService)
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}

	// Listen for shutdown signal and gracefully shutdown gRPC and NATS
	shutdown.ListenForShutdown(shutdown.ShutdownConfig{
		GrpcServer:      grpcServerInstance,
		NATSConnection:  natsConn,
		ShutdownTimeout: 5 * time.Second,
	})
}
