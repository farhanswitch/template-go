package models

import (
	"time"

	"github.com/google/uuid"
)

type Author struct {
	ID         uuid.UUID `json:"uuid"`
	Name       string    `json:"name"`
	Created    time.Time `json:"-"`
	Updated    time.Time `json:"-"`
	StrCreated string    `json:"created"`
	StrUpdated string    `json:"updated"`
}

type CreateAuthorRequest struct {
	UUID string `json:"uuid"`
	Name string `json:"name" validate:"required,min=3,max=100"`
}
