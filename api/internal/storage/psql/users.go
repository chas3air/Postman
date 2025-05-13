package psql

import (
	"api/internal/domain/models"
	storageerrors "api/internal/storage"
	"api/pkg/lib/logger/sl"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

const UsersTableName = "users"

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
	const op = "storage.GetUsers"
	log := p.log.With(
		"op", op,
	)

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("%s: %w", op, ctx.Err())
	default:
	}

	rows, err := p.DB.QueryContext(ctx, `
	SELECT id, login, password FROM `+UsersTableName+`;
	`)
	if err != nil {
		log.Error("Cannot scanning rows", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var users = make([]models.User, 0, 10)
	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Id, &user.Login, &user.Password); err != nil {
			log.Warn("Error scanning row", sl.Err(err))
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

// GetUserById implements storage.IUsersStorage.
func (p *PsqlStorage) GetUserById(ctx context.Context, id uuid.UUID) (models.User, error) {
	const op = "storage.GetUserById"
	log := p.log.With(
		"op", op,
	)

	select {
	case <-ctx.Done():
		return models.User{}, fmt.Errorf("%s: %w", op, ctx.Err())
	default:
	}

	var user models.User
	err := p.DB.QueryRowContext(ctx, `
		SELECT id, login, password FROM `+UsersTableName+`
		WHERE id = $1
	`, id).Scan(&user.Id, &user.Login, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Error("User with current id not found", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, storageerrors.ErrNotFound)
		}

		log.Error("Error scanning row", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// InsertUser implements storage.IUsersStorage.
func (p *PsqlStorage) InsertUser(ctx context.Context, user models.User) (models.User, error) {
	const op = "storage.InsertUser"
	log := p.log.With(
		"op", op,
	)

	select {
	case <-ctx.Done():
		return models.User{}, fmt.Errorf("%s: %w", op, ctx.Err())
	default:
	}

	_, err := p.DB.ExecContext(ctx, `
		INSERT INTO `+UsersTableName+`
		VALUES ($1, $2, $3);
	`, user.Id, user.Login, user.Password)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			log.Error("User with this id already exists", sl.Err(err))
			return models.User{}, fmt.Errorf("%s: %w", op, storageerrors.ErrAlreadyExists)
		}

		log.Error("Error inserting user", sl.Err(err))
		return models.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}

// UpdateUser implements storage.IUsersStorage.
func (p *PsqlStorage) UpdateUser(ctx context.Context, id uuid.UUID, user models.User) (models.User, error) {
	panic("unimplemented")
}

// DeleteUSer implements storage.IUsersStorage.
func (p *PsqlStorage) DeleteUser(ctx context.Context, id uuid.UUID) (models.User, error) {
	panic("unimplemented")
}
