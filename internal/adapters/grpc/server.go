package grpc

import (
	"github.com/bohexists/task-manager-svc/ports/inbound"
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/bohexists/task-manager-svc/api/proto"
)

// TaskService implements pb.TaskServiceServer
type TaskService struct {
	pb.UnimplementedTaskServiceServer
	service inbound.TaskServiceServer
}

// StartGRPCServer starts gRPC server
func StartGRPCServer(service *inbound.TaskServiceServer) *grpc.Server {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, TaskService{})

	go func() {
		log.Println("gRPC server is running on port :50051")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	return s
}
