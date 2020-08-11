package domain

import "time"

// Todo model
type Todo struct {
	ID        int
	Title     string
	Content   string
	Deadline  time.Time
	State     int
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Todos Todoの複数形
type Todos []Todo
