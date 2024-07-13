package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var db *sql.DB

func init() {
	var err error

	db, err = sql.Open("sqlite3", "/tasks.db")

	if err != nil {
		log.Fatal(err)
	}

	createTable()
}

func createTable() {
	createTableSQL := `CREATE TABLE IF NOT EXIST tasks (
    	"id" TEXT NOT NULL PRIMARY KEY,
    	"title" TEXT,
    	"description" TEXT,
    	"completed" BOOLEAN
    );`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, title, description, completed FROM tasks")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []Task

	for rows.Next() {
		var task Task

		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)

		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func getTaskByID(id string) (*Task, error) {
	row := db.QueryRow("SELECT  id, title, description, completed FROM tasks WHERE id = ?", id)
	var task Task

	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Completed)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func addTask(task *Task) error {
	_, err := db.Exec("INSERT INTO tasks (id, title, decription, completed) VALUES (?,?, ?, ?)",
		task.ID, task.Title, task.Description, task.Completed)
	return err
}

func updateTaskByID(task *Task) error {
	_, err := db.Exec("UPDATE tasks SET title = ?, description = ?, completed = ? where id = ?",
		task.Title, task.Description, task.Completed, task.ID)

	return err
}

func deleteTaskByID(id string) error {
	_, err := db.Exec("DELETE FROM TASK WHERE ID = ?", id)

	return err
}
