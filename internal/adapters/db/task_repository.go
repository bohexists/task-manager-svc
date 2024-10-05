package db

import (
	"database/sql"
	"github.com/bohexists/task-manager-svc/api/proto"
)

type TaskRepository struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *proto.Task) (int64, error) {
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

func (r *TaskRepository) GetTask(id int64) (*proto.Task, error) {
	var task proto.Task
	err := r.DB.QueryRow("SELECT id, title, description FROM tasks WHERE id = ?", id).Scan(&task.Id, &task.Title, &task.Description)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) UpdateTask(task *proto.Task) error {
	_, err := r.DB.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", task.Title, task.Description, task.Id)
	return err
}

func (r *TaskRepository) DeleteTask(id int64) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (r *TaskRepository) ListTasks() ([]*proto.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, description FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*proto.Task
	for rows.Next() {
		var task proto.Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description); err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	return tasks, nil
}
