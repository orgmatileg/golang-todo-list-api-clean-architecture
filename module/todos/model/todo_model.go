package model

import (
	"time"
)

type Todo struct {
	TodoID    string    `json:"todo_id"`
	TodoName  string    `json:"todo_name"`
	IsDone    bool      `json:"is_done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Todos []Todo

func NewTodo() *Todo {
	return &Todo{
		IsDone:    false,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
