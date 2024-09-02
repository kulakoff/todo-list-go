package repositories

import (
	"fmt"
	"github.com/kulakoff/todo-list-go/cmd/models"
	"github.com/kulakoff/todo-list-go/cmd/storage"
	"log"
)

func CreateTask(task models.Task) (models.Task, error) {
	log.Println("DB || CreateTask")
	db := storage.GetDB()
	sqlQuerry := `INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) returning id`
	err := db.QueryRow(sqlQuerry, task.Title, task.Description, task.DueDate).Scan(&task.ID)
	if err != nil {
		log.Println(err.Error())
		log.Println(task)
		return task, err
	}

	return task, nil
}

func GetTask(id int) (models.Task, error) {
	log.Println("DB || GetTask")

	db := storage.GetDB()
	sqlQuerry := `SELECT * FROM tasks WHERE id = $1`
	row := db.QueryRow(sqlQuerry, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		log.Println("SQL err")
		log.Println(err.Error())
		return models.Task{}, err
	}
	return task, nil
}

func GetAllTasks() ([]models.Task, error) {
	log.Println("DB || GetAllTasks")

	var tasks []models.Task
	db := storage.GetDB()
	sqlQuery := "SELECT id, title, description, due_date, created_at, updated_at FROM tasks ORDER BY id"
	rows, err := db.Query(sqlQuery)
	if err != nil {
		log.Println("SQL ERR:")
		log.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
			log.Println("Error scanning row:")
			log.Println(err.Error())
			continue
		}
		tasks = append(tasks, task)
	}

	// Проверка на наличие ошибок после завершения цикла
	if err := rows.Err(); err != nil {
		log.Println("Rows iteration error:")
		log.Println(err.Error())
		return nil, err
	}

	return tasks, nil
}

func UpdateTask(task models.Task, id int) (models.Task, error) {
	log.Println("DB || UpdateTask")

	db := storage.GetDB()
	sqlQuerry := `UPDATE tasks
	SET title = $2, description = $3, due_date = $4, updated_at = $5
	WHERE id = $1
	RETURNING id`

	err := db.QueryRow(sqlQuerry, id, task.Title, task.Description, task.DueDate, task.UpdatedAt).Scan(&task.ID)
	if err != nil {
		log.Println("SQL ERR:")
		log.Println(err.Error())
		return task, err
	}
	task.ID = id
	return task, nil
}

func DeleteTask(id int) error {
	log.Println("DB || DeleteTask")

	db := storage.GetDB()
	sqlQuerry := `DELETE FROM tasks WHERE id = $1`
	result, err := db.Exec(sqlQuerry, id)
	if err != nil {
		log.Println("SQL ERR:")
		log.Println(err.Error())
		//return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:")
		return err
	}

	if rowsAffected == 0 {
		log.Println("No rows were affected")
		return fmt.Errorf("not found ID: %d", id)
	}

	return nil
}
