package models

import "github.com/google/uuid"

type User struct {
	ID          uuid.UUID
	Username    string
	Password    string
	AccessLevel bool // ?
}
