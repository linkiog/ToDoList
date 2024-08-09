package backend

import "time"

type Task struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	DueDate   time.Time `json:"due_date"`
	Priority  int       `json:"priority"`
}
