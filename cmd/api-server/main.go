package main

import (
	"dhturdav/golang-todo-list/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	repo := repository.NewInMemory()

	router.GET("/todos", func(c *gin.Context) {
		c.JSON(http.StatusOK, repo.List())
	})

	router.Run()
}
