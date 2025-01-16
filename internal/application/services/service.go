package services

import (
	"github.com/Jereyji/avito/internal/domain/entities"
	"github.com/Jereyji/avito/internal/domain/interface_repository"
	"golang.org/x/net/context"
)

type EstateService struct {
	repository interface_repository.RepositoryInterface
}

func NewEstateService(rep interface_repository.RepositoryInterface) *EstateService {
	return &EstateService{
		repository: rep,
	}
}

func (s EstateService) Register(ctx context.Context, username string, password string, accessLevel int) error {
	user, err := entities.NewUser(username, password, accessLevel)
	if err != nil {
		return err
	}

	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return err // ADD: error handling?
	}

	return nil
}

func (s EstateService) Login(ctx context.Context, username, password string) (*entities.Tokens, error) {
	user, err := s.repository.FetchUser(ctx, username)
	if err != nil {
		return nil, err
	}

	err = user.VerifyPassword(password)
	if err != nil {
		return nil, err
	}

	// TODO: add generate tokens with user.AccessLevel

	return nil, nil
}

func (s EstateService) DummyLogin(ctx context.Context, username string, password string) (*entities.User, error) {
	return nil, nil
}

func (s EstateService) GetFlatsByHouseId(houseId string) ([]entities.Flat, error) {
	return nil, nil
}
