package main

import (
	"github.com/kulakoff/todo-list-go/internal/app/endpoint"
	"github.com/kulakoff/todo-list-go/internal/app/service"
	"github.com/kulakoff/todo-list-go/internal/repositories"
	"github.com/kulakoff/todo-list-go/internal/storage"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	//	---- DB connect ----
	db, err := storage.New()
	if err != nil {
		panic(err)
	}

	// ----- Repository & Service init -----
	taskRepo := repositories.New(db)
	taskService := service.New(taskRepo)
	handler := endpoint.New(taskService)

	e := echo.New()

	// ---- Routes ----
	e.GET("/tasks", handler.GetAll).Name = "get-all"
	e.POST("/tasks", handler.Create).Name = "create"
	e.GET("/tasks/:id", handler.Get).Name = "get"
	e.PUT("/tasks/:id", handler.Update).Name = "update"
	e.DELETE("/tasks/:id", handler.Delete).Name = "delete"

	//	---- Start server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
