package repository

import "dhturdav/golang-todo-list/internal/entity"

type InMemory struct {
	todos []entity.Todo
}

func NewInMemory() *InMemory {
	return &InMemory{}
}

func (r *InMemory) List() []entity.Todo {
	return r.todos
}
