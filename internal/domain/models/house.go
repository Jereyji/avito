package models

import (
	"github.com/google/uuid"
	"time"
)

type House struct {
	ID               uuid.UUID
	Street           string
	Builder          string
	YearConstruction time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
