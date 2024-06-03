package entity

import "github.com/google/uuid"

type Todo struct {
	ID    uuid.UUID `json:"id"`
	Title string    `json:"title"`
}
