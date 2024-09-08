package service

import (
	"database/sql"
	"errors"
	"github.com/kulakoff/todo-list-go/internal/err_msg"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"log/slog"
)

type TaskService interface {
	GetAllTasks() ([]repositories.Task, error)
	GetTask(id int) (repositories.Task, error)
	CreateTask(task repositories.Task) (repositories.Task, error)
	UpdateTask(id int, task repositories.Task) (repositories.Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo repositories.TaskRepository
}

func New(repo repositories.TaskRepository) *taskService {
	return &taskService{repo: repo}
}

func (t *taskService) GetAllTasks() ([]repositories.Task, error) {
	tasks, err := t.repo.GetAllTasks()
	if err != nil {
		slog.Error("Failed to get all tasks", "error", err)
		return nil, err
	}
	return tasks, nil
}

func (t *taskService) CreateTask(task repositories.Task) (repositories.Task, error) {
	if err := task.Validate(); err != nil {
		slog.Warn("Task validation failed", "error", err)
		return repositories.Task{}, err
	}

	createdTask, err := t.repo.CreateTask(task)
	if err != nil {
		return repositories.Task{}, err
	}
	return createdTask, nil
}

func (t *taskService) GetTask(id int) (repositories.Task, error) {
	task, err := t.repo.GetTaskById(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			slog.Warn("Task not found", "taskID", id)
			return repositories.Task{}, err_msg.ErrTaskNotFound
		}
		return repositories.Task{}, err
	}
	return task, nil
}

func (t *taskService) UpdateTask(id int, task repositories.Task) (repositories.Task, error) {
	if err := task.Validate(); err != nil {
		slog.Warn("Task validation failed", "error", err)
		return repositories.Task{}, err
	}

	updatedTask, err := t.repo.UpdateTask(id, task)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			slog.Warn("Task not found for update", "taskID", id)
			return repositories.Task{}, err_msg.ErrTaskNotFound
		}
		slog.Error("Failed to update task", "error", err)
		return repositories.Task{}, err
	}
	return updatedTask, nil
}

func (t *taskService) DeleteTask(id int) error {
	err := t.repo.DeleteTask(id)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			slog.Warn("Task not found for deletion", "taskID", id)
			return err_msg.ErrTaskNotFound
		}
		slog.Error("Failed to delete task", "error", err)
		return err
	}
	return nil
}
