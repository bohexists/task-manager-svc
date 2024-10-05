package grpc

import (
	"context"
	"log"

	pb "github.com/bohexists/task-manager-svc/api/proto"
)

// TaskService реализует gRPC сервис для управления задачами
type TaskService struct {
	pb.UnimplementedTaskServiceServer
}

// CreateTask создает новую задачу
func (s *TaskService) CreateTask(ctx context.Context, req *pb.Task) (*pb.TaskID, error) {
	// Логика для создания задачи (вставка в базу данных)
	log.Printf("Creating new task: %v", req)

	// Для демонстрации возвращаем сгенерированный ID задачи
	return &pb.TaskID{Id: req.Id}, nil
}

// GetTask получает задачу по ID
func (s *TaskService) GetTask(ctx context.Context, req *pb.TaskID) (*pb.Task, error) {
	// Логика для получения задачи из базы данных по ID
	log.Printf("Fetching task with ID: %d", req.Id)

	// Для демонстрации возвращаем фиктивную задачу
	return &pb.Task{
		Id:          req.Id,
		Title:       "Sample Task",
		Description: "This is a sample task",
	}, nil
}

// UpdateTask обновляет задачу
func (s *TaskService) UpdateTask(ctx context.Context, req *pb.Task) (*pb.Empty, error) {
	// Логика для обновления задачи в базе данных
	log.Printf("Updating task with ID: %d", req.Id)

	// Возвращаем пустой ответ
	return &pb.Empty{}, nil
}

// DeleteTask удаляет задачу по ID
func (s *TaskService) DeleteTask(ctx context.Context, req *pb.TaskID) (*pb.Empty, error) {
	// Логика для удаления задачи из базы данных
	log.Printf("Deleting task with ID: %d", req.Id)

	// Возвращаем пустой ответ
	return &pb.Empty{}, nil
}

// ListTasks возвращает поток задач
func (s *TaskService) ListTasks(req *pb.Empty, stream pb.TaskService_ListTasksServer) error {
	// Логика для выборки всех задач из базы данных и отправки их клиенту через stream
	log.Println("Listing all tasks")

	// Для демонстрации отправляем фиктивные задачи
	for i := int64(1); i <= 3; i++ {
		task := &pb.Task{
			Id:          i,
			Title:       "Task " + string(i),
			Description: "Description for task " + string(i),
		}
		if err := stream.Send(task); err != nil {
			return err
		}
	}
	return nil
}
