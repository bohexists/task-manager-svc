package app

import (
	"context"
	"github.com/bohexists/task-manager-svc/domain"
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

func (s *TaskService) CreateTask(ctx context.Context, title, description string) (int64, error) {
	// Используем бизнес-логику домена
	task, err := domain.NewTask(title, description)
	if err != nil {
		return 0, err
	}
	return s.TaskRepo.CreateTask(task)
}

func (s *TaskService) GetTask(ctx context.Context, id int64) (*domain.Task, error) {
	return s.TaskRepo.GetTask(id)
}

func (s *TaskService) UpdateTask(ctx context.Context, id int64, title, description string) error {
	task, err := s.TaskRepo.GetTask(id)
	if err != nil {
		return err
	}

	// Обновляем задачу через бизнес-логику домена
	if err := task.Update(title, description); err != nil {
		return err
	}

	return s.TaskRepo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(ctx context.Context, id int64) error {
	return s.TaskRepo.DeleteTask(id)
}

func (s *TaskService) ListTasks(ctx context.Context) ([]*domain.Task, error) {
	return s.TaskRepo.ListTasks()
}
