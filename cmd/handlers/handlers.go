package handlers

import (
	"errors"
	"github.com/kulakoff/todo-list-go/cmd/models"
	"github.com/kulakoff/todo-list-go/cmd/repositories"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"strconv"
	"time"
)

func GetAll(c echo.Context) error {
	log.Println("run GetAll")

	tasks, err := repositories.GetAllTasks()
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, tasks)
}

func Get(c echo.Context) error {
	log.Println("run Get task")
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	task, err := repositories.GetTask(idInt)
	if err != nil {
		if errors.Is(err, repositories.ErrTaskNotFound) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, task)
}

func CreateTask(c echo.Context) error {
	// TODO: Implement check payload data
	task := models.Task{}
	err := c.Bind(&task)
	if err != nil {
		log.Println("Failed bind data")
		return err
	}
	task.CreatedAt = time.Now()
	task.UpdatedAt = task.CreatedAt

	newTask, err := repositories.CreateTask(task)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newTask)
}

func UpdateTask(c echo.Context) error {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": err.Error()})
	}

	task := models.Task{}
	task.UpdatedAt = time.Now()

	err = c.Bind(&task)
	if err != nil {
		return err
	}

	updatedTask, err := repositories.UpdateTask(task, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, updatedTask)
}

func DeleteTask(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Error converting id to int")
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Invalid task ID "})
	}

	err = repositories.DeleteTask(idInt)
	if err != nil {
		if errors.Is(err, repositories.ErrTaskNotFound) {
			log.Println("Error deleting task, not found")
			return c.JSON(http.StatusNotFound, map[string]string{"error": "task not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
	}

	return c.JSON(http.StatusNoContent, nil)
}
