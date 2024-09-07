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

func New() (*sql.DB, error) {
	if db != nil {
		return db, nil
	}

	// ----- load config -----
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	// ----- check connection
	err = connect(config)
	if err != nil {
		return nil, err
	}

	// ----- run migration
	err = migrate()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadConfig() (map[string]string, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config := map[string]string{
		"host":     os.Getenv("POSTGRES_HOST"),
		"port":     os.Getenv("POSTGRES_PORT"),
		"user":     os.Getenv("POSTGRES_USER"),
		"password": os.Getenv("POSTGRES_PASSWORD"),
		"dbname":   os.Getenv("POSTGRES_DB"),
	}

	for key, value := range config {
		if value == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", key)
		}
	}
	return config, nil
}

func connect(config map[string]string) error {
	var err error
	db, err = sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			config["host"], config["user"], config["password"], config["dbname"], config["port"]))
	if err != nil {
		return fmt.Errorf("failed to open database connection: %w", err)
	}

	// ----- check connection
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("failed to ping database connection: %w", err)
	}

	slog.Info("Successfully connected to DB")
	return nil
}

func migrate() error {
	var exists bool
	// ----- check table exist
	err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'tasks')").Scan(&exists)
	if err != nil {
		return err
	}

	// ----- run migration if table not exist
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
