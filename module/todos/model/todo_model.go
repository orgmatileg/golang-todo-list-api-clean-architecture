package model

import (
	"time"
)

// Todo Struct
type Todo struct {
	TodoID    string    `json:"todo_id"`
	TodoName  string    `json:"todo_name"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Todo list
type Todos []Todo

// NewTodo func
func NewTodo() *Todo {
	return &Todo{
		IsDone:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
