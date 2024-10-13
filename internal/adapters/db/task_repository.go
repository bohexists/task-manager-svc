package db

import (
	"database/sql"
	"github.com/bohexists/task-manager-svc/domain"
	"github.com/bohexists/task-manager-svc/ports/outbound"
)

// TaskRepository implements TaskRepository interface now works with domain.Task instead of proto.Task
type TaskRepository struct {
	DB *sql.DB
}

// NewTaskRepository creates a new TaskRepository
func NewTaskRepository(db *sql.DB) outbound.TaskRepository {
	return &TaskRepository{DB: db}
}

// CreateTask creates a new task
func (r *TaskRepository) CreateTask(task *domain.Task) (int64, error) {
	result, err := r.DB.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
	if err != nil {
		return 0, err
	}
	taskID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return taskID, nil
}

// GetTask gets a task
func (r *TaskRepository) GetTask(id int64) (*domain.Task, error) {
	var task domain.Task
	err := r.DB.QueryRow("SELECT id, title, description, status FROM tasks WHERE id = ?", id).Scan(&task.ID, &task.Title, &task.Status)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

// UpdateTask updates a task
func (r *TaskRepository) UpdateTask(task *domain.Task) error {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", task.Title, task.Description, task.ID)
	return err
}

// DeleteTask deletes a task
func (r *TaskRepository) DeleteTask(id int64) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

// ListTasks lists all tasks
func (r *TaskRepository) ListTasks() ([]*domain.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, description, status FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*domain.Task
	for rows.Next() {
		var task domain.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}

// UpdateTaskStatus updates only the task's status
func (r *TaskRepository) UpdateTaskStatus(id int64, status string) error {
	_, err := r.DB.Exec("UPDATE tasks SET status = ? WHERE id = ?", status, id)
	return err
}
