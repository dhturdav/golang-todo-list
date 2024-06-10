package api_test

import (
	"bytes"
	"dhturdav/golang-todo-list/internal/api"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
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

func TestPostTodos(t *testing.T) {
	router := api.Router()
	todo := map[string]string{"title": "Test todo"}
	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "Expected POST /todos to succeed")
}

func TestPostTodosResponseIDIsUUID(t *testing.T) {
	router := api.Router()
	todo := map[string]string{"title": "Test todo"}
	jsonValue, _ := json.Marshal(todo)
	req, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	_, err := uuid.Parse(response["id"].(string))
	assert.Nil(t, err, "Expected response ID to be a valid UUID")
}

func TestAddTwoTodosThenGet(t *testing.T) {
	router := api.Router()
	todo1 := map[string]string{"title": "First todo"}
	todo2 := map[string]string{"title": "Second todo"}
	jsonValue1, _ := json.Marshal(todo1)
	jsonValue2, _ := json.Marshal(todo2)

	req1, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue1))
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	req2, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonValue2))
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	req3, _ := http.NewRequest("GET", "/todos", nil)
	w3 := httptest.NewRecorder()
	router.ServeHTTP(w3, req3)

	var response1, response2 map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &response1)
	json.Unmarshal(w2.Body.Bytes(), &response2)

	id1, _ := response1["id"].(string)
	id2, _ := response2["id"].(string)

	expectedResponse := `[{"id":"` + id1 + `","title":"First todo"},{"id":"` + id2 + `","title":"Second todo"}]`
	assert.JSONEq(t, expectedResponse, w3.Body.String(), "Expected GET /todos to return the list of added todos")
}
