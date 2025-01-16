package interface_repository

import (
	"context"
	"github.com/Jereyji/avito/internal/domain/entities"
)

type RepositoryInterface interface {
	FetchUser(ctx context.Context, username string) (*entities.User, error)
	CreateUser(ctx context.Context, user *entities.User) error
	UpdateUser(ctx context.Context, user *entities.User) (*entities.User, error)
	DeleteUser(ctx context.Context, userId string) error

	FetchHouse(ctx context.Context, houseId string) (*entities.House, error)
	CreateHouse(ctx context.Context, house *entities.House) error
	UpdateHouse(ctx context.Context, house *entities.House) (*entities.House, error)
	DeleteHouse(ctx context.Context, houseId string) error

	FetchFlat(ctx context.Context, flatId string) (*entities.Flat, error)
	CreateFlat(ctx context.Context, flat *entities.Flat) error
	UpdateFlat(ctx context.Context, flat *entities.Flat) (*entities.Flat, error)
	DeleteFlat(ctx context.Context, flatId string) error
}

//HouseSubscribe(houseId string) (error)
