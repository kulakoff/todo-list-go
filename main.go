package main

import (
	"github.com/kulakoff/todo-list-go/cmd/handlers"
	"github.com/kulakoff/todo-list-go/cmd/storage"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	// ---- Routes ----
	e.POST("/tasks", handlers.CreateTask)
	e.GET("/tasks", handlers.GetAll)
	e.PUT("/tasks/:id", handlers.UpdateTask)
	e.GET("/tasks/:id", handlers.Get)
	e.DELETE("/tasks/:id", handlers.DeleteTask)

	//	---- DB connect ----
	storage.InitDB()

	//	---- Start server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
