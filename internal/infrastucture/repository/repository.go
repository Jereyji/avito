package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"sync"
)

type EstateRepository struct {
	db *pgxpool.Pool
	mu sync.RWMutex
}

func NewEstateRepository(db *pgxpool.Pool) *EstateRepository {
	return &EstateRepository{
		db: db,
	}
}
