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

	r.Create(todo)

	todos := r.List()
	assert.Equal(t, 1, len(todos), "expected 1 todo after creation")
	assert.Equal(t, "Test todo", todos[0].Title, "expected todo title to match")
	assert.NotEqual(t, uuid.Nil, todos[0].ID, "expected todo to have a non-nil ID")
}
