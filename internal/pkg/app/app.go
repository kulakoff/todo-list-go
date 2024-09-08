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
	Handler    *endpoint.Endpoint
	Service    *service.TaskService
	Repository *repositories.TaskRepository
}

func New() (*App, error) {
	// ----- init storage
	db, err := storage.New()
	if err != nil {
		panic(err)
	}

	a := &App{}
	a.Echo = echo.New()
	a.Repository = repositories.New(db)

	//// ----- init repository, service and endpoint-handlers
	//taskRepo := repositories.New(db)
	//taskService := service.New(taskRepo)
	//handler := endpoint.New(taskService)
	//
	//e := echo.New()
	//
	//return &App{
	//	Echo:       e,
	//	Handler:    handler,
	//	Service:    taskService,
	//	Repository: taskRepo,
	//}, nil
}

func (a *App) SetupRoutes() {
	// ---- Routes ----
	a.Echo.GET("/tasks", a.Handler.GetAll).Name = "get-all"
	a.Echo.POST("/tasks", a.Handler.Create).Name = "create"
	a.Echo.GET("/tasks/:id", a.Handler.Get).Name = "get"
	a.Echo.PUT("/tasks/:id", a.Handler.Update).Name = "update"
	a.Echo.DELETE("/tasks/:id", a.Handler.Delete).Name = "delete"
}
