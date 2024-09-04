package storage

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"os"
)

var db *sql.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbHost, dbUser, dbPass, dbName, dbPort))
	if err != nil {
		slog.Error(err.Error())
	}

	err = db.Ping()
	if err != nil {
		slog.Error(err.Error())
	}

	slog.Info("Successfully connect to DB")

	// ----- run migration
	err = migrate()
	if err != nil {
		slog.Error("Error running migration: ", err)
	}
}

func New() *sql.DB {
	return db
}

func migrate() error {
	var exists bool
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'tasks')").Scan(&exists)
	if err != nil {
		return err
	}

	// run migration if table not exist
	if !exists {
		sqlCreateTable := `
		CREATE TABLE tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT,
			due_date TIMESTAMP NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
		);`

		_, err := db.Exec(sqlCreateTable)
		if err != nil {
			return err
		}

		slog.Info("Migration completed successfully: Table created")
	} else {
		slog.Info("Migration completed successfully: Table already exists")
	}
	return nil
}
