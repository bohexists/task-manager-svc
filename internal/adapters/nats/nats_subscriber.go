package nats

import (
	"encoding/json"
	"github.com/bohexists/task-manager-svc/ports/outbound"
	"github.com/nats-io/nats.go"
	"log"

	"github.com/bohexists/task-manager-svc/config"
)

// TaskStatusUpdate starts with capital letter
type TaskStatusUpdate struct {
	TaskID int64  `json:"task_id"`
	Status string `json:"status"`
}

// InitNATSSubscriber initializes NATS subscriber
func InitNATSSubscriber(cfg config.Config, repo outbound.TaskRepository) (*nats.Conn, error) {
	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		return nil, err
	}

	subject := "task.status.update"
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		var update TaskStatusUpdate
		err := json.Unmarshal(m.Data, &update)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		// Обновление статуса задачи в базе данных
		err = repo.UpdateTaskStatus(update.TaskID, update.Status)
		if err != nil {
			log.Printf("Error updating task status in DB: %v", err)
			return
		}

		log.Printf("Task ID %d status updated to %s", update.TaskID, update.Status)
	})

	if err != nil {
		return nil, err
	}

	log.Printf("Subscribed to NATS subject %s", subject)
	return nc, nil
}

// SubscribeToTaskStatusUpdates connects to NATS and subscribes to a subject
func SubscribeToTaskStatusUpdates(cfg config.Config, repo outbound.TaskRepository) error {
	// Connect to NATS
	nc, err := nats.Connect(cfg.NatsURL)
	if err != nil {
		return err
	}
	defer nc.Close()

	// Subscribe to NATS subject
	subject := "task.status.update"
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		var update TaskStatusUpdate
		err := json.Unmarshal(m.Data, &update)
		if err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			return
		}

		// Update task status in the database
		err = repo.UpdateTaskStatus(update.TaskID, update.Status)
		if err != nil {
			log.Printf("Error updating task status in DB: %v", err)
			return
		}

		log.Printf("Task ID %d status updated to %s", update.TaskID, update.Status)
	})

	if err != nil {
		return err
	}

	log.Printf("Subscribed to NATS subject %s", subject)
	return nil
}
