package models

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}
