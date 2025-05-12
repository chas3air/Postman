package usersservice

import (
	"api/internal/domain/interfaces/storage"
	"api/internal/domain/models"
	"context"
	"log/slog"

	"github.com/google/uuid"
)

type UsersService struct {
	log     *slog.Logger
	storage *storage.IUsersStorage
}

func New(log *slog.Logger, storage storage.IUsersStorage) *UsersService {
	return &UsersService{
		log:     log,
		storage: &storage,
	}
}

// GetUsers implements service.IUsersService.
func (u *UsersService) GetUsers(ctx context.Context) ([]models.User, error) {
	panic("unimplemented")
}

// GetUserById implements service.IUsersService.
func (u *UsersService) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

// InsertUser implements service.IUsersService.
func (u *UsersService) InsertUser(ctx context.Context, user models.User) (models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements service.IUsersService.
func (u *UsersService) UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error) {
	panic("unimplemented")
}

// DeleteUSer implements service.IUsersService.
func (u *UsersService) DeleteUSer(ctx context.Context, id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}
