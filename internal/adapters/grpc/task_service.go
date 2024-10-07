package grpc

import (
	"context"
	"github.com/bohexists/task-manager-svc/api/proto"
	"github.com/bohexists/task-manager-svc/app"
)

type TaskServiceServer struct {
	proto.UnimplementedTaskServiceServer
	TaskService *app.TaskService
}

func NewTaskServiceServer(taskService *app.TaskService) *TaskServiceServer {
	return &TaskServiceServer{
		TaskService: taskService,
	}
}

func (s *TaskServiceServer) CreateTask(ctx context.Context, req *proto.Task) (*proto.TaskID, error) {
	return s.TaskService.CreateTask(ctx, req)
}

func (s *TaskServiceServer) GetTask(ctx context.Context, req *proto.TaskID) (*proto.Task, error) {
	return s.TaskService.GetTask(ctx, req)
}

func (s *TaskServiceServer) UpdateTask(ctx context.Context, req *proto.Task) (*proto.Empty, error) {
	return s.TaskService.UpdateTask(ctx, req)
}

func (s *TaskServiceServer) DeleteTask(ctx context.Context, req *proto.TaskID) (*proto.Empty, error) {
	return s.TaskService.DeleteTask(ctx, req)
}

func (s *TaskServiceServer) ListTasks(req *proto.Empty, stream proto.TaskService_ListTasksServer) error {
	return s.TaskService.ListTasks(stream.Context(), req, stream)
}
