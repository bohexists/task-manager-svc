package domain

import "errors"

const (
	StatusNew        = "new"
	StatusInProgress = "in_progress"
	StatusCompleted  = "completed"
)

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      string
}

// NewTask creates a new task with default status as "new"
func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	// Create a new task with default status as "new"
	return &Task{
		Title:       title,
		Description: description,
		Status:      StatusNew,
	}, nil
}

// Update updates a task's title and description
func (t *Task) Update(title, description string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	t.Title = title
	t.Description = description
	return nil
}

// UpdateStatus updates a task's status
func (t *Task) UpdateStatus(status string) error {
	if status != StatusNew && status != StatusInProgress && status != StatusCompleted {
		return errors.New("invalid status")
	}
	t.Status = status
	return nil
}
