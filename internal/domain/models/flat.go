package models

import "github.com/google/uuid"

type Flat struct {
	ID              uuid.UUID
	ApartmentNumber int
	Price           int
	CountRooms      int
}
