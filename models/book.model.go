package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID       uuid.UUID `json:"id"`
	AuthorID uuid.UUID `json:"authorID"`
	Title    string    `json:"title"`
	Amount   uint16    `json:"amount"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
