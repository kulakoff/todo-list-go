package repositories

import "database/sql"

type TaskRepository interface {
	CreateTask(task Task) (Task, error)
	GetAllTasks() ([]Task, error)
	GetTaskById(id int) (Task, error)
	UpdateTask(task Task, id int) (Task, error)
	DeleteTask(id int) error
}

type taskRepository struct {
	db *sql.DB
}

func (t taskRepository) CreateTask(task Task) (Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) GetAllTasks() ([]Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) GetTaskById(id int) (Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) UpdateTask(task Task, id int) (Task, error) {
	//TODO implement me
	panic("implement me")
}

func (t taskRepository) DeleteTask(id int) error {
	//TODO implement me
	panic("implement me")
}

func New(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}
