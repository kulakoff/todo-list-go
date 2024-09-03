package main

import (
	"github.com/kulakoff/todo-list-go/internal/handlers"
	"github.com/kulakoff/todo-list-go/internal/storage"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	taskHandler := handlers.NewTaskHandler()

	// ---- Routes ----
	e.GET("/tasks", taskHandler.GetAll)
	e.POST("/tasks", taskHandler.Create)
	e.GET("/tasks/:id", taskHandler.Get)
	e.PUT("/tasks/:id", taskHandler.Update)
	e.DELETE("/tasks/:id", taskHandler.Delete)

	//	---- DB connect ----
	storage.InitDB()

	//	---- Start server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
