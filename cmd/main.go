package main

import (
	"github.com/bohexists/task-manager-svc/app"
	"github.com/bohexists/task-manager-svc/config"
	"github.com/bohexists/task-manager-svc/internal/adapters/db"
	"github.com/bohexists/task-manager-svc/internal/adapters/grpc"
)

func main() {
	cfg := config.LoadConfig()

	db.ConnectToDB(cfg)

	taskRepo := db.NewTaskRepository(db.DB)

	taskService := app.NewTaskService(taskRepo)

	grpcServiceServer := grpc.NewTaskServiceServer(taskService)
	grpc.StartGRPCServer(grpcServiceServer)
}
