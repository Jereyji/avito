package interface_repository

import "github.com/Jereyji/avito/internal/domain/models"

type RepositoryInterface interface {
	FetchUser(username string) (*models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(user *models.User) (*models.User, error)
	DeleteUser(user *models.User) error

	FetchHouse(houseId string) (*models.House, error)
	CreateHouse(house *models.House) (*models.House, error)
	UpdateHouse(house *models.House) (*models.House, error)
	DeleteHouse(houseId string) error

	FetchFlat(flatId string) (*models.Flat, error)
	CreateFlat(flat *models.Flat) (*models.Flat, error)
	UpdateFlat(flat *models.Flat) (*models.Flat, error)
	DeleteFlat(flat *models.Flat) error
}

//HouseSubscribe(houseId string) (error)
