package repositories

import (
	"database/sql"
	"errors"
	"github.com/kulakoff/todo-list-go/cmd/err_msg"
	"github.com/kulakoff/todo-list-go/cmd/models"
	"github.com/kulakoff/todo-list-go/cmd/storage"
	"log"
)

//var ErrTaskNotFound = err_msg.New("task not found")

func CreateTask(task models.Task) (models.Task, error) {
	db := storage.GetDB()
	sqlQuerry := `INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) returning id`
	err := db.QueryRow(sqlQuerry, task.Title, task.Description, task.DueDate).Scan(&task.ID)
	if err != nil {
		return task, err
	}

	return task, nil
}

func GetTask(id int) (models.Task, error) {
	db := storage.GetDB()
	sqlQuerry := `SELECT * FROM tasks WHERE id = $1`
	row := db.QueryRow(sqlQuerry, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, err_msg.ErrTaskNotFound
		}
		return models.Task{}, err
	}
	return task, nil
}

func GetAllTasks() ([]models.Task, error) {
	db := storage.GetDB()
	sqlQuery := "SELECT id, title, description, due_date, created_at, updated_at FROM tasks ORDER BY id"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		return nil, err
	}

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			log.Println("Error scanning row:")
			log.Println(err.Error())
			continue
		}
		tasks = append(tasks, task)
	}

	// Checking for err_msg after a cycle has completed
	if err := rows.Err(); err != nil {
		log.Println("Rows iteration error:")
		log.Println(err.Error())
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(task models.Task, id int) (models.Task, error) {
	db := storage.GetDB()
	sqlQuerry := `UPDATE tasks
	SET title = $2, description = $3, due_date = $4, updated_at = $5
	WHERE id = $1
	RETURNING id`

	err := db.QueryRow(sqlQuerry, id, task.Title, task.Description, task.DueDate, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return task, err_msg.ErrTaskNotFound
		}

		return task, err
	}
	task.ID = id

	return task, nil
}

func DeleteTask(id int) error {
	db := storage.GetDB()
	sqlQuerry := `DELETE FROM tasks WHERE id = $1`
	result, err := db.Exec(sqlQuerry, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:")
		return err
	}

	if rowsAffected == 0 {
		log.Println("No rows were affected")
		return err_msg.ErrTaskNotFound
	}

	return nil
}
