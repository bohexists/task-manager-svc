package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/bohexists/task-manager-svc/api/proto"
)

type server struct {
	pb.UnimplementedTaskServiceServer
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTaskServiceServer(s, &server{})
	log.Printf("gRPC server is running on port :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
