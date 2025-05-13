package usersservice

import (
	"api/internal/domain/interfaces/storage"
	"api/internal/domain/models"
	serviceerrors "api/internal/service"
	storageerrors "api/internal/storage"
	"api/pkg/lib/logger/sl"
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
)

type UsersService struct {
	log     *slog.Logger
	storage storage.IUsersStorage
}

func New(log *slog.Logger, storage storage.IUsersStorage) *UsersService {
	return &UsersService{
		log:     log,
		storage: storage,
	}
}

// GetUsers implements service.IUsersService.
func (u *UsersService) GetUsers(ctx context.Context) ([]models.User, error) {
	const op = "services.GetUsers"
	log := u.log.With(
		"op", op,
	)

	users, err := u.storage.GetUsers(ctx)
	if err != nil {
		if errors.Is(err, storageerrors.ErrNotFound) {
			log.Warn("No one user in storage", sl.Err(serviceerrors.ErrNotFound))
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		log.Error("Error fetching users", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}

// GetUserById implements service.IUsersService.
func (u *UsersService) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	const op = "services.GetUserById"
	log := u.log.With(
		"op", op,
	)

	user, err := u.storage.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, storageerrors.ErrNotFound) {
			log.Warn("No one user with id", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, serviceerrors.ErrNotFound)
		}

		log.Error("Error fetching user by id", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// InsertUser implements service.IUsersService.
func (u *UsersService) InsertUser(ctx context.Context, user models.User) (models.User, error) {
	const op = "services.InsertUser"
	log := u.log.With(
		"op", op,
	)

	user, err := u.storage.InsertUser(ctx, user)
	if err != nil {
		if errors.Is(err, storageerrors.ErrAlreadyExists) {
			log.Warn("User already exists", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, serviceerrors.ErrAlreadyExists)
		}

		log.Error("Error inserting user", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// UpdateUser implements service.IUsersService.
func (u *UsersService) UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error) {
	const op = "services.UpdateUser"
	log := u.log.With(
		"op", op,
	)

	user, err := u.storage.UpdateUser(ctx, id, user)
	if err != nil {
		if errors.Is(err, storageerrors.ErrNotFound) {
			log.Warn("No one user with current id", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, serviceerrors.ErrNotFound)
		}

		log.Error("Error updating user", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// DeleteUSer implements service.IUsersService.
func (u *UsersService) DeleteUSer(ctx context.Context, id uuid.UUID) (models.User, error) {
	const op = "services.DeleteUser"
	log := u.log.With(
		"op", op,
	)

	user, err := u.storage.DeleteUser(ctx, id)
	if err != nil {
		if errors.Is(err, storageerrors.ErrNotFound) {
			log.Warn("No one user with current id", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, serviceerrors.ErrNotFound)
		}

		log.Error("Error deleting user", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
