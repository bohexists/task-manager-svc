package grpc

import (
	"context"
	"github.com/bohexists/task-manager-svc/internal/adapters/db"
	"log"

	pb "github.com/bohexists/task-manager-svc/api/proto"
)

// TaskService
type TaskService struct {
	pb.UnimplementedTaskServiceServer
	TaskRepo *db.TaskRepository
}

func NewTaskService(repo *db.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req *pb.Task) (*pb.TaskID, error) {
	log.Printf("Creating new task: %v", req)

	taskID, err := s.TaskRepo.CreateTask(req)
	if err != nil {
		log.Printf("Failed to create task: %v", err)
		return nil, err
	}

	return &pb.TaskID{Id: taskID}, nil
}

func (s *TaskService) GetTask(ctx context.Context, req *pb.TaskID) (*pb.Task, error) {
	log.Printf("Fetching task with ID: %d", req.Id)

	task, err := s.TaskRepo.GetTask(req.Id)
	if err != nil {
		log.Printf("Failed to fetch task: %v", err)
		return nil, err
	}
	return task, nil
}

func (s *TaskService) UpdateTask(ctx context.Context, req *pb.Task) (*pb.Empty, error) {
	log.Printf("Updating task with ID: %d", req.Id)

	err := s.TaskRepo.UpdateTask(req)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return nil, err
	}

	return &pb.Empty{}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *pb.TaskID) (*pb.Empty, error) {
	log.Printf("Deleting task with ID: %d", req.Id)

	err := s.TaskRepo.DeleteTask(req.Id)
	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *TaskService) ListTasks(req *pb.Empty, stream pb.TaskService_ListTasksServer) error {
	log.Println("Listing all tasks")

	tasks, err := s.TaskRepo.ListTasks()
	if err != nil {
		log.Printf("Failed to list tasks: %v", err)
		return err
	}

	for _, task := range tasks {
		if err := stream.Send(task); err != nil {
			return err
		}
	}

	return nil
}
