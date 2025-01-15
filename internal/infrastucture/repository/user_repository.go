package repository

import (
	"context"
	"errors"
	"github.com/Jereyji/avito/internal/domain/models"
	"github.com/Jereyji/avito/internal/infrastucture/repository/queries"
	"github.com/jackc/pgx/v5"
)

func (r *EstateRepository) FetchUser(ctx context.Context, username string) (*models.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var user models.User
	err := r.db.QueryRow(ctx, queries.FetchUserByUsernameQuery, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.AccessLevel,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &user, nil
		}

		return nil, err
	}

	return &user, nil
}

func (r *EstateRepository) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.CreateUserQuery,
		user.ID,
		user.Username,
		user.Password,
		user.AccessLevel,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *EstateRepository) UpdateUser(ctx context.Context, user *models.User) (*models.User, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.UpdateUserQuery,
		user.ID,
		user.Username,
		user.Password,
		user.AccessLevel,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *EstateRepository) DeleteUser(ctx context.Context, userId string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, err := r.db.Exec(ctx, queries.DeleteUserQuery, userId)
	if err != nil {
		return err
	}

	return nil
}
