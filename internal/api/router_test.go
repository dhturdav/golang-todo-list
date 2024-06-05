package api_test

import (
	"dhturdav/golang-todo-list/internal/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := api.Router()

	assert.NotNil(t, router, "Expected router to not be nil")
}

func TestTodosRouteRegistered(t *testing.T) {
	router := api.Router()
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected /todos route to be registered")
}

func TestTodosRouteReturnsEmptyList(t *testing.T) {
	router := api.Router()
	req, _ := http.NewRequest("GET", "/todos", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, "[]", w.Body.String(), "Expected /todos route to return an empty list")
}
