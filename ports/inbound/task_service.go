package inbound

import (
	"context"
	"github.com/bohexists/task-manager-svc/api/proto"
	"github.com/bohexists/task-manager-svc/app"
)

// TaskServiceServer implements proto.TaskServiceServer
type TaskServiceServer struct {
	TaskService *app.TaskService
}

// NewTaskServiceServer creates new TaskServiceServer
func NewTaskServiceServer(taskService *app.TaskService) *TaskServiceServer {
	return &TaskServiceServer{
		TaskService: taskService,
	}
}

// CreateTask implements proto.TaskServiceServer
func (s *TaskServiceServer) CreateTask(ctx context.Context, req *proto.Task) (*proto.TaskID, error) {
	taskID, err := s.TaskService.CreateTask(ctx, req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	return &proto.TaskID{Id: taskID}, nil
}

// GetTask implements proto.TaskServiceServer
func (s *TaskServiceServer) GetTask(ctx context.Context, req *proto.TaskID) (*proto.Task, error) {
	task, err := s.TaskService.GetTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &proto.Task{
		Id:          task.ID,
		Title:       task.Title,
		Description: task.Description,
	}, nil
}

// UpdateTask implements proto.TaskServiceServer
func (s *TaskServiceServer) UpdateTask(ctx context.Context, req *proto.Task) (*proto.Empty, error) {
	err := s.TaskService.UpdateTask(ctx, req.Id, req.Title, req.Description)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

// DeleteTask implements proto.TaskServiceServer
func (s *TaskServiceServer) DeleteTask(ctx context.Context, req *proto.TaskID) (*proto.Empty, error) {
	err := s.TaskService.DeleteTask(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

// ListTasks implements proto.TaskServiceServer
func (s *TaskServiceServer) ListTasks(req *proto.Empty, stream proto.TaskService_ListTasksServer) error {
	tasks, err := s.TaskService.ListTasks(stream.Context())
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if err := stream.Send(&proto.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
		}); err != nil {
			return err
		}
	}

	return nil
}
