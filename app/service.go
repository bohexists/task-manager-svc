package app

import (
	"context"
	"github.com/bohexists/task-manager-svc/api/proto"
	"github.com/bohexists/task-manager-svc/ports/outbound"
)

type TaskService struct {
	TaskRepo outbound.TaskRepository
}

func NewTaskService(repo outbound.TaskRepository) *TaskService {
	return &TaskService{
		TaskRepo: repo,
	}
}

func (s *TaskService) CreateTask(ctx context.Context, req *proto.Task) (*proto.TaskID, error) {
	taskID, err := s.TaskRepo.CreateTask(req)
	if err != nil {
		return nil, err
	}
	return &proto.TaskID{Id: taskID}, nil
}

func (s *TaskService) GetTask(ctx context.Context, req *proto.TaskID) (*proto.Task, error) {
	return s.TaskRepo.GetTask(req.Id)
}

func (s *TaskService) UpdateTask(ctx context.Context, req *proto.Task) (*proto.Empty, error) {
	err := s.TaskRepo.UpdateTask(req)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (s *TaskService) DeleteTask(ctx context.Context, req *proto.TaskID) (*proto.Empty, error) {
	err := s.TaskRepo.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &proto.Empty{}, nil
}

func (s *TaskService) ListTasks(ctx context.Context, req *proto.Empty, stream proto.TaskService_ListTasksServer) error {
	tasks, err := s.TaskRepo.ListTasks()
	if err != nil {
		return err
	}

	for _, task := range tasks {
		if err := stream.Send(task); err != nil {
			return err
		}
	}
	return nil
}
