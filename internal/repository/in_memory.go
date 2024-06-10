package repository

import (
	"dhturdav/golang-todo-list/internal/entity"
	"github.com/google/uuid"
)

type InMemory struct {
	todos []entity.Todo
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (r *InMemory) List() []entity.Todo {
	if len(r.todos) == 0 {
		return []entity.Todo{}
	}

	return r.todos
}

func (r *InMemory) Create(todo entity.Todo) entity.Todo {
	todo.ID = uuid.New()
	r.todos = append(r.todos, todo)

	return todo
}
