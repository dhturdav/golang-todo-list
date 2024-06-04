package api

import (
	"dhturdav/golang-todo-list/internal/repository"

	"github.com/gin-gonic/gin"
)

type todosHandler struct {
	repository *repository.InMemory
}

func (t *todosHandler) list(c *gin.Context) {
	c.JSON(200, t.repository.List())
}
