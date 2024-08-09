package backend

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	*sql.DB
}

func NewDB() *DB {
	db, err := sql.Open("sqlite3", "./tasks.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		completed BOOLEAN,
		due_date DATETIME,
		priority INTEGER
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

func (db *DB) AddTask(name string, dueDate time.Time, priority int) error {
	_, err := db.Exec("INSERT INTO tasks (name, completed, due_date, priority) VALUES (?, ?, ?, ?)", name, false, dueDate, priority)
	return err
}

func (db *DB) DeleteTask(id int) error {
	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	return err
}

func (db *DB) GetTasks() ([]Task, error) {
	rows, err := db.Query("SELECT id, name, completed, priority, due_date FROM tasks ORDER BY priority ASC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.ID, &task.Name, &task.Completed, &task.Priority, &task.DueDate); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (db *DB) MarkTaskCompleted(id int) error {
	_, err := db.Exec("UPDATE tasks SET completed = 1 WHERE id = ?", id)
	return err
}

func (db *DB) UndoTaskCompletion(id int) error {
	_, err := db.Exec("UPDATE tasks SET completed = false WHERE id = ?", id)
	return err
}
