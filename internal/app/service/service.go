package service

import (
	"database/sql"
	"errors"
	"github.com/kulakoff/todo-list-go/internal/err_msg"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"log/slog"
	"time"
)

type TaskService interface {
	CreateTask(task repositories.Task) (repositories.Task, error)
	GetTask(id int) (repositories.Task, error)
	GetAllTasks() ([]repositories.Task, error)
	UpdateTask(id int, task repositories.Task) (repositories.Task, error)
	DeleteTask(id int) error
}

type TaskServiceStruct struct {
	repo repositories.TaskRepository
}

func (t *TaskServiceStruct) CreateTask(task repositories.Task) (repositories.Task, error) {
	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	createdTask, err := t.repo.CreateTask(task)
	if err != nil {
		return repositories.Task{}, err
	}
	return createdTask, nil
}

func (t *TaskServiceStruct) GetTask(id int) (repositories.Task, error) {
	task, err := t.repo.GetTaskById(id)
	if err != nil {
		slog.Info("service.GetTask ", err.Error())
		if errors.Is(err, sql.ErrNoRows) {
			slog.Info(err.Error())
			return task, err_msg.ErrTaskNotFound
		}
		return task, err
	}
	return task, nil
}

func (t *TaskServiceStruct) GetAllTasks() ([]repositories.Task, error) {
	tasks, err := t.repo.GetAllTasks()
	if err != nil {
		return []repositories.Task{}, err
	}
	return tasks, nil
}

func (t *TaskServiceStruct) UpdateTask(id int, task repositories.Task) (repositories.Task, error) {
	task = repositories.Task{}
	task.UpdatedAt = time.Now()

	updatedTask, err := t.repo.UpdateTask(id, task)
	if err != nil {
		return repositories.Task{}, err
	}
	return updatedTask, nil
}

func (t *TaskServiceStruct) DeleteTask(id int) error {
	err := t.repo.DeleteTask(id)
	if err != nil {
		return err
	}
	return nil
}

func New(repo repositories.TaskRepository) *TaskServiceStruct {
	return &TaskServiceStruct{repo: repo}

}
