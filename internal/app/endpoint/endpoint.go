package endpoint

import (
	"errors"
	"github.com/kulakoff/todo-list-go/internal/err_msg"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"github.com/labstack/echo/v4"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

//type IntEndpoint interface {
//	GetAll(c echo.Context) error
//	Get(c echo.Context) error
//	Create(c echo.Context) error
//	Update(c echo.Context) error
//	Delete(c echo.Context) error
//}

type Service interface {
}

type Endpoint struct {
	s Service
}

func (e *Endpoint) GetAll(c echo.Context) error {
	tasks, err := repositories.GetAllTasks()
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal)
	}
	return c.JSON(http.StatusOK, tasks)
}

func (e *Endpoint) Get(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Info("failed parse ID to int")
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err_msg.ErrBadRequest.Error()})
	}

	task, err := repositories.GetTask(idInt)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, err_msg.ErrTaskNotFound)
		}
		return c.JSON(http.StatusInternalServerError, err_msg.ErrInternal)
	}

	return c.JSON(http.StatusOK, task)
}

func (e *Endpoint) Create(c echo.Context) error {
	// TODO: Implement check payload data
	task := models.Task{}
	err := c.Bind(&task)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err_msg.ErrBadRequest)
	}

	now := time.Now()
	task.CreatedAt = now
	task.UpdatedAt = now

	newTask, err := repositories.CreateTask(task)
	if err != nil {
		slog.Info(err.Error())
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Проблема на сервере"})
	}
	return c.JSON(http.StatusCreated, newTask)
}

func (e *Endpoint) Update(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err_msg.ErrBadRequest)
	}

	task := models.Task{}
	task.UpdatedAt = time.Now()

	err = c.Bind(&task)
	if err != nil {
		return err
	}

	updatedTask, err := repositories.UpdateTask(task, idInt)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}

		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func (e *Endpoint) Delete(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		slog.Info("Error converting id to int")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid task ID "})
	}

	err = repositories.DeleteTask(idInt)
	if err != nil {
		if errors.Is(err, err_msg.ErrTaskNotFound) {
			slog.Info("Error deleting task, not found")
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}

func New() *Endpoint {
	return &Endpoint{}
}
