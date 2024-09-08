package endpoint

import (
	"errors"
	"github.com/kulakoff/todo-list-go/internal/app/service"
	"github.com/kulakoff/todo-list-go/internal/err_msg"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
)

type Endpoint interface {
	GetAll(c echo.Context) error
	Get(c echo.Context) error
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
}

type endpoint struct {
	s service.TaskService
}

func (e *endpoint) GetAll(c echo.Context) error {
	tasks, err := e.s.GetAllTasks()
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal)
	}
	return c.JSON(http.StatusOK, tasks)
}

func (e *endpoint) Get(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Info("failed parse ID to int")
		return c.JSON(http.StatusBadRequest, err_msg.ErrBadRequest.Error())
	}

	task, err := e.s.GetTask(idInt)
	if err != nil {
		slog.Info("endpoint.get", err.Error())
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, err_msg.ErrTaskNotFound.Error())
		}
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func (e *endpoint) Create(c echo.Context) error {
	// TODO: Implement check payload data
	task := repositories.Task{}
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err_msg.ErrBadRequest.Error())
	}

	newTask, err := e.s.CreateTask(task)
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal.Error())
	}
	return c.JSON(http.StatusCreated, newTask)
}

func (e *endpoint) Update(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err_msg.ErrBadRequest.Error())
	}

	task := repositories.Task{}

	err = c.Bind(&task)
	if err != nil {
		return err
	}

	updatedTask, err := e.s.UpdateTask(idInt, task)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, err_msg.ErrTaskNotFound.Error())
		}

		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal.Error())
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (e *endpoint) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Info("Error converting id to int")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid task ID "})
	}

	err = e.s.DeleteTask(idInt)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, err_msg.ErrTaskNotFound.Error())
		}
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal.Error())
	}

	return c.JSON(http.StatusNoContent, nil)
}

func New(s service.TaskService) *endpoint {
	return &endpoint{s: s}
}
