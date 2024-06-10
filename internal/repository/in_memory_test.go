package repository_test

import (
	"dhturdav/golang-todo-list/internal/entity"
	"dhturdav/golang-todo-list/internal/repository"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestList(t *testing.T) {
	r := repository.NewInMemory()

	todos := r.List()

	if len(todos) != 0 {
		t.Errorf("expected 0 todos, got %d", len(todos))
	}

	if todos == nil {
		t.Errorf("expected todos not to be nil")
	}
}

func TestCreate(t *testing.T) {
	r := repository.NewInMemory()
	todo := entity.Todo{Title: "Test todo"}

	createdTodo := r.Create(todo)
	assert.Equal(t, "Test todo", createdTodo.Title, "expected todo title to match")
	assert.NotEqual(t, uuid.Nil, createdTodo.ID, "expected todo to have a non-nil ID")
}
