package domain

import "errors"

type Task struct {
	ID          int64
	Title       string
	Description string
}

// NewTask creates a new task
func NewTask(title, description string) (*Task, error) {
	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	return &Task{
		Title:       title,
		Description: description,
	}, nil
}

// Update updates a task
func (t *Task) Update(title, description string) error {
	if title == "" {
		return errors.New("title cannot be empty")
	}
	t.Title = title
	t.Description = description
	return nil
}
