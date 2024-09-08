package main

import (
	"github.com/kulakoff/todo-list-go/internal/pkg/app"
	"log/slog"
	"os"
)

func main() {
	a, err := app.New()
	if err != nil {
		slog.Error(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5055"
	}
	err = a.Start(port)
	if err != nil {
		slog.Error(err.Error())
	}
}
