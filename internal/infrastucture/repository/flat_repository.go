package repository

import (
	"context"
	"github.com/Jereyji/avito/internal/domain/entities"
	"github.com/Jereyji/avito/internal/infrastucture/repository/queries"
)

func (r *EstateRepository) FetchFlat(ctx context.Context, flatId string) (*entities.Flat, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var flat entities.Flat
	err := r.db.QueryRow(ctx, queries.FetchFlatQuery, flatId).Scan(
		flat.ID,
		flat.HouseId,
		flat.Price,
		flat.CountRooms,
		flat.ApartmentNumber,
		flat.ModerationStatus,
	)
	if err != nil {
		return nil, err
	}

	return &flat, nil
}

func (r *EstateRepository) CreateFlat(ctx context.Context, flat *entities.Flat) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.CreateFlatQuery,
		flat.ID,
		flat.HouseId,
		flat.Price,
		flat.CountRooms,
		flat.ApartmentNumber,
		flat.ModerationStatus,
	)
	if err != nil {
		return err
	}

	return nil
}

func (r *EstateRepository) UpdateFlat(ctx context.Context, flat *entities.Flat) (*entities.Flat, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.UpdateFlatQuery,
		flat.ID,
		flat.HouseId,
		flat.Price,
		flat.CountRooms,
		flat.ApartmentNumber,
		flat.ModerationStatus,
	)
	if err != nil {
		return nil, err
	}

	return flat, err
}

func (r *EstateRepository) DeleteFlat(ctx context.Context, flatId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.DeleteFlatQuery, flatId)
	if err != nil {
		return err
	}

	return nil
}
