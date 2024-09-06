package repositories

//import (
//	"database/sql"
//	"errors"
//	"github.com/kulakoff/todo-list-go/internal/err_msg"
//	"github.com/kulakoff/todo-list-go/internal/storage"
//	"log/slog"
//)
//
//func CreateTask(task Task) (Task, error) {
//	db := storage.New()
//	sqlQuerry := `INSERT INTO tasks (title, description, due_date) VALUES ($1, $2, $3) returning id`
//	err := db.QueryRow(sqlQuerry, task.Title, task.Description, task.DueDate).Scan(&task.ID)
//	if err != nil {
//		return task, err
//	}
//
//	return task, nil
//}
//
//func GetTask(id int) (Task, error) {
//	db := storage.New()
//	sqlQuerry := `SELECT * FROM tasks WHERE id = $1`
//	row := db.QueryRow(sqlQuerry, id)
//
//	var task Task
//	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return task, err_msg.ErrTaskNotFound
//		}
//		return Task{}, err
//	}
//	return task, nil
//}
//
//func GetAllTasks() ([]Task, error) {
//	db := storage.New()
//	sqlQuery := "SELECT id, title, description, due_date, created_at, updated_at FROM tasks ORDER BY id"
//	rows, err := db.Query(sqlQuery)
//	if err != nil {
//		return nil, err
//	}
//
//	var tasks []Task
//	for rows.Next() {
//		var task Task
//		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.CreatedAt, &task.UpdatedAt); err != nil {
//			slog.Info("Error scanning row:", err)
//			continue
//		}
//		tasks = append(tasks, task)
//	}
//
//	// Checking for err_msg after a cycle has completed
//	if err := rows.Err(); err != nil {
//		slog.Info("Rows iteration error:", err)
//		return nil, err
//	}
//
//	return tasks, nil
//}
//
//func UpdateTask(task Task, id int) (Task, error) {
//	db := storage.New()
//	sqlQuerry := `UPDATE tasks
//	SET title = $2, description = $3, due_date = $4, updated_at = $5
//	WHERE id = $1
//	RETURNING id`
//
//	err := db.QueryRow(sqlQuerry, id, task.Title, task.Description, task.DueDate, task.UpdatedAt).Scan(&task.ID)
//	if err != nil {
//		if errors.Is(err, sql.ErrNoRows) {
//			return task, err_msg.ErrTaskNotFound
//		}
//
//		return task, err
//	}
//	task.ID = id
//
//	return task, nil
//}
//
//func DeleteTask(id int) error {
//	db := storage.New()
//	sqlQuerry := `DELETE FROM tasks WHERE id = $1`
//	result, err := db.Exec(sqlQuerry, id)
//	if err != nil {
//		return err
//	}
//
//	rowsAffected, err := result.RowsAffected()
//	if err != nil {
//		slog.Info("Error checking rows affected:")
//		return err
//	}
//
//	if rowsAffected == 0 {
//		slog.Info("No rows were affected")
//		return err_msg.ErrTaskNotFound
//	}
//
//	return nil
//}
