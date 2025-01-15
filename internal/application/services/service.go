package services

import (
	"github.com/Jereyji/avito/internal/domain/interface_repository"
	"github.com/Jereyji/avito/internal/domain/models"
)

type EstateService struct {
	repository interface_repository.RepositoryInterface
}

func NewEstateService(rep interface_repository.RepositoryInterface) *EstateService {
	return &EstateService{
		repository: rep,
	}
}

func (s EstateService) Register(user *models.User) (*models.User, error) {
	// TODO: add hashing user's password

	return s.repository.CreateUser(user)
}

func (s EstateService) Login(username, password string) (*models.Tokens, error) {
	user, err := s.repository.User(username)
	if err != nil {
		return nil, err
	}

	// TODO: add compare password hashes
	//err := compareHashPassword(password, user.Password)

	// TODO: add generate tokens with user.AccessLevel

	return _, nil
}

func (s EstateService) DummyLogin(username string, password string) (*models.User, error) {
}

func (s EstateService) GetFlatsByHouseId(houseId string) ([]models.Flat, error) {

}
