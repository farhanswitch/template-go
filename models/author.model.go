package models

import "github.com/google/uuid"

type Author struct {
	ID   uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type CreateAuthorRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}
