package repository_test

import (
	"dhturdav/golang-todo-list/internal/repository"
	"testing"
)

func TestList(t *testing.T) {
	r := repository.NewInMemory()

	todos := r.List()

	if len(todos) != 0 {
		t.Errorf("expected 0 todos, got %d", len(todos))
	}
}
