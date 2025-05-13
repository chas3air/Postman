package storage

import (
	"api/internal/domain/models"
	"context"

	"github.com/google/uuid"
)

type IUsersStorage interface {
	GetUsers(ctx context.Context) ([]models.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (models.User, error)
	InsertUser(ctx context.Context, user models.User) (models.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (models.User, error)
}
