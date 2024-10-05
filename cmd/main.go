package main

import (
	"github.com/bohexists/task-manager-svc/config"
	"github.com/bohexists/task-manager-svc/internal/adapters/db"
	"github.com/bohexists/task-manager-svc/internal/adapters/grpc"
)

func main() {
	cfg := config.LoadConfig()
	db.ConnectToDB(cfg)
	grpc.StartGRPCServer()
}
