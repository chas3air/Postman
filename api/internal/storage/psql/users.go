package psql

import (
	"api/internal/domain/models"
	"api/pkg/lib/logger/sl"
	"context"
	"database/sql"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

type PsqlStorage struct {
	log *slog.Logger
	DB  *sql.DB
}

func New(log *slog.Logger, connStr string) *PsqlStorage {
	const op = "psql.New"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.With("op", op).Error("Error connecting to DB", sl.Err(err))
		panic(err)
	}

	wd, _ := os.Getwd()
	migrationPath := filepath.Join(wd, "app", "migrations")
	if err := applyMigrations(db, migrationPath); err != nil {
		panic(err)
	}

	return &PsqlStorage{
		log: log,
		DB:  db,
	}
}

func applyMigrations(db *sql.DB, migrationsPath string) error {
	return goose.Up(db, migrationsPath)
}

func (p *PsqlStorage) Close() {
	p.DB.Close()
}

// GetUsers implements storage.IUsersStorage.
func (p *PsqlStorage) GetUsers(ctx context.Context) ([]models.User, error) {
	panic("unimplemented")
}

// GetUserById implements storage.IUsersStorage.
func (p *PsqlStorage) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}

// InsertUser implements storage.IUsersStorage.
func (p *PsqlStorage) InsertUser(ctx context.Context, user models.User) (models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements storage.IUsersStorage.
func (p *PsqlStorage) UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error) {
	panic("unimplemented")
}

// DeleteUSer implements storage.IUsersStorage.
func (p *PsqlStorage) DeleteUSer(ctx context.Context, id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}
