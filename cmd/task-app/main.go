package main

import (
	"github.com/kulakoff/todo-list-go/internal/app/endpoint"
	"github.com/kulakoff/todo-list-go/internal/storage"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	handler := endpoint.New()

	// ---- Routes ----
	e.GET("/tasks", handler.GetAll).Name = "get-all"
	e.POST("/tasks", handler.Create).Name = "create"
	e.GET("/tasks/:id", handler.Get).Name = "get"
	e.PUT("/tasks/:id", handler.Update).Name = "update"
	e.DELETE("/tasks/:id", handler.Delete).Name = "delete"

	//	---- DB connect ----
	storage.InitDB()

	//	---- Start server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
