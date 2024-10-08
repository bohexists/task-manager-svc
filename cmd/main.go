package main

import (
	"github.com/bohexists/task-manager-svc/app"
	"github.com/bohexists/task-manager-svc/config"
	"github.com/bohexists/task-manager-svc/internal/adapters/db"
	"github.com/bohexists/task-manager-svc/internal/adapters/grpc"
	"github.com/bohexists/task-manager-svc/ports/inbound"
)

func main() {
	// Load config
	cfg := config.LoadConfig()

	// Connect to database
	db.ConnectToDB(cfg)

	// Initialize repository, service, and handler
	taskRepo := db.NewTaskRepository(db.DB)
	// Initialize service
	taskService := app.NewTaskService(taskRepo)

	// Initialize grpc server
	grpcServiceServer := inbound.NewTaskServiceServer(taskService)
	// Start grpc server
	grpc.StartGRPCServer(grpcServiceServer)

}
