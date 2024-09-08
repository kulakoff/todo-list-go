package app

import (
	"github.com/kulakoff/todo-list-go/internal/app/endpoint"
	"github.com/kulakoff/todo-list-go/internal/app/service"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"github.com/kulakoff/todo-list-go/internal/storage"
	"github.com/labstack/echo/v4"
)

type App struct {
	Echo       *echo.Echo
	Handler    endpoint.Endpoint
	Service    service.TaskService
	Repository repositories.TaskRepository
}

func New() (*App, error) {
	// ----- init storage
	db, err := storage.New()
	if err != nil {
		panic(err)
	}

	taskRepo := repositories.New(db)
	taskService := service.New(taskRepo)
	handler := endpoint.New(taskService)

	var a = &App{
		Echo:       echo.New(),
		Handler:    handler,
		Service:    taskService,
		Repository: taskRepo,
	}
	return a, nil
}

func (a *App) SetupRoutes() {
	a.Echo.GET("/tasks", a.Handler.GetAll).Name = "get-all"
	a.Echo.POST("/tasks", a.Handler.Create).Name = "create"
	a.Echo.GET("/tasks/:id", a.Handler.Get).Name = "get"
	a.Echo.PUT("/tasks/:id", a.Handler.Update).Name = "update"
	a.Echo.DELETE("/tasks/:id", a.Handler.Delete).Name = "delete"
}

func (a *App) Start(port string) error {
	if port == "" {
		port = "5055"
	}
	return a.Echo.Start(":" + port)
}
