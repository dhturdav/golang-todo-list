package repository

import "dhturdav/golang-todo-list/internal/entity"

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
