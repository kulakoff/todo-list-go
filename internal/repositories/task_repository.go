package repositories

import (
	"database/sql"
	"errors"
	"github.com/kulakoff/todo-list-go/internal/err_msg"
	"log/slog"
)

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id int) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
	DeleteTask(id int) error
}

type taskRepository struct {
	db *sql.DB
}

const (
	selectTaskById = `SELECT id, title, description, due_date, created_at, updated_at FROM tasks WHERE id = $1`
	selectAllTasks = `SELECT id, title, description, due_date, created_at, updated_at FROM tasks ORDER BY id`
	insertTask     = `INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) RETURNING id, title, description, due_date, created_at, updated_at`
	updateTask     = `UPDATE tasks SET title = $2, description = $3, due_date = $4, updated_at = $5 WHERE id = $1 RETURNING id, title, description, due_date, created_at, updated_at`
	deleteTask     = `DELETE FROM tasks WHERE id = $1`
)

func (t *taskRepository) CreateTask(task Task) (Task, error) {
	err := t.db.QueryRow(insertTask, task.Title, task.Description, task.DueDate).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		slog.Error("Failed to create task", err)
		return task, err
	}
	return task, nil
}

func (t *taskRepository) GetAllTasks() ([]Task, error) {
	rows, err := t.db.Query(selectAllTasks)
	if err != nil {
		slog.Error("Failed to get all tasks", err)
		return nil, err
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			slog.Info("Error scanning row:", err)
			continue
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		slog.Info("Rows iteration error:", err)
		return nil, err
	}

	return tasks, nil
}

func (t *taskRepository) GetTaskById(id int) (Task, error) {
	var task Task
	err := t.db.QueryRow(selectTaskById, id).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		//if errors.Is(err, sql.ErrNoRows) {
		//	slog.Info(err.Error())
		//	return task, err_msg.ErrTaskNotFound
		//}
		return task, err
	}
	return task, nil
}

func (t *taskRepository) UpdateTask(id int, task Task) (Task, error) {
	err := t.db.QueryRow(updateTask, id, task.Title, task.Description, task.DueDate, task.UpdatedAt).Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, err_msg.ErrTaskNotFound
		}
		return task, err
	}
	return task, nil
}

func (t *taskRepository) DeleteTask(id int) error {
	result, err := t.db.Exec(deleteTask, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Info("Error checking rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		slog.Info("No rows were affected", "taskID", id)
		return err_msg.ErrTaskNotFound
	}

	return nil
}

func New(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}
