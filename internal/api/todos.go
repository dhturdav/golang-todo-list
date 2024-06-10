package api

import (
	"dhturdav/golang-todo-list/internal/entity"
	"dhturdav/golang-todo-list/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todosHandler struct {
	repository *repository.InMemory
}

func (t *todosHandler) list(c *gin.Context) {
	c.JSON(200, t.repository.List())
}

func (t *todosHandler) create(c *gin.Context) {
	var todo entity.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todo = t.repository.Create(todo)
	c.JSON(http.StatusOK, todo)
}
