package main

import "dhturdav/golang-todo-list/internal/api"

func main() {
	r := api.Router()

	r.Run()
}
