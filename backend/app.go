package backend

import (
	"time"
)

type App struct {
	db *DB
}

func NewApp() *App {
	db := NewDB()
	return &App{db: db}
}

func (a *App) AddTask(name string, dueDate time.Time, priority int) error {
	return a.db.AddTask(name, dueDate, priority)
}

func (a *App) DeleteTask(id int) error {
	return a.db.DeleteTask(id)
}

func (a *App) GetTasks() ([]Task, error) {
	return a.db.GetTasks()
}

func (a *App) MarkTaskCompleted(id int) error {
	return a.db.MarkTaskCompleted(id)
}

func (a *App) UndoTaskCompletion(id int) error {
	return a.db.UndoTaskCompletion(id)
}
