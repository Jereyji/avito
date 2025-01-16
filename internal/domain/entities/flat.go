package entities

import "github.com/google/uuid"

type Flat struct {
	ID               uuid.UUID
	HouseId          uuid.UUID
	Price            int
	CountRooms       int
	ApartmentNumber  int
	ModerationStatus int
}
