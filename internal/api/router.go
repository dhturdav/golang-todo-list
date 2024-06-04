package api

import (
	"dhturdav/golang-todo-list/internal/repository"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()
	todos := todosHandler{repository: repository.NewInMemory()}

	router.GET("/todos", todos.list)

	return router
}
