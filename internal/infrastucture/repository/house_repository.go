package repository

import (
	"context"
	"github.com/Jereyji/avito/internal/domain/models"
	"github.com/Jereyji/avito/internal/infrastucture/repository/queries"
)

func (r *EstateRepository) FetchHouse(ctx context.Context, houseId string) (*models.House, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var house models.House
	err := r.db.QueryRow(ctx, queries.FetchHouseQuery, houseId).Scan(
		&house.ID,
		&house.Street,
		&house.Builder,
		&house.YearConstruction,
		&house.CreatedAt,
		&house.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &house, nil
}

func (r *EstateRepository) CreateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.CreateHouseQuery,
		house.ID,
		house.Street,
		house.Builder,
		house.YearConstruction,
		house.CreatedAt,
		house.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (r *EstateRepository) UpdateHouse(ctx context.Context, house *models.House) (*models.House, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.UpdateHouseQuery,
		house.ID,
		house.Street,
		house.Builder,
		house.YearConstruction,
		house.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return house, nil
}

func (r *EstateRepository) DeleteHouse(ctx context.Context, houseId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.DeleteHouseQuery, houseId)
	if err != nil {
		return err
	}

	return nil
}
