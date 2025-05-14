package models

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID `json:"id,omitempty"`
	Login    string    `json:"login,omitempty"`
	Password string    `json:"password,omitempty"`
}
