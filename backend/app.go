package backend

import (
	"errors"
	"fmt"
	"time"
)

type App struct {
	db *DB
}

func NewApp() *App {
	db := NewDB()
	return &App{db: db}
}

func ValidateTaskName(name string) error {
	if len(name) == 0 {
		return errors.New("task name cannot be empty")
	}
	if len(name) > 255 {
		return errors.New("task name is too long")
	}
	return nil
}

func ValidatePriority(priority int) error {
	if priority < 1 || priority > 5 {
		return errors.New("priority must be between 1 and 5")
	}
	return nil
}

func (a *App) AddTask(name string, dueDate time.Time, priority int) error {
	if err := ValidateTaskName(name); err != nil {
		return fmt.Errorf("task validation failed: %w", err)
	}
	if err := ValidatePriority(priority); err != nil {
		return fmt.Errorf("priority validation failed: %w", err)
	}

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
