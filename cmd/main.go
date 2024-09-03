package main

import (
	"github.com/kulakoff/todo-list-go/internal/handlers"
	"github.com/kulakoff/todo-list-go/internal/storage"
	"github.com/labstack/echo/v4"
	"os"
)

func main() {
	e := echo.New()

	// ---- Routes ----
	e.GET("/tasks", handlers.GetAll)
	e.POST("/tasks", handlers.Create)
	e.GET("/tasks/:id", handlers.Get)
	e.PUT("/tasks/:id", handlers.Update)
	e.DELETE("/tasks/:id", handlers.Delete)

	//	---- DB connect ----
	storage.InitDB()

	//	---- Start server ----
	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	e.Logger.Fatal(e.Start(":" + port))
}
