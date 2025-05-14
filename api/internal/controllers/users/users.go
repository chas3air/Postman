package userscontroller

import (
	"api/internal/domain/interfaces/service"
	"api/internal/domain/models"
	serviceerrors "api/internal/service"
	"api/pkg/lib/logger/sl"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UsersController struct {
	log     *slog.Logger
	service service.IUsersService
}

func New(log *slog.Logger, service service.IUsersService) *UsersController {
	return &UsersController{
		log:     log,
		service: service,
	}
}

func (u *UsersController) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.GetUsersHandler"
	log := u.log.With(
		"op", op,
	)

	users, err := u.service.GetUsers(r.Context())
	if err != nil {
		if errors.Is(err, serviceerrors.ErrNotFound) {
			log.Warn("No users in storage", sl.Err(err))
			json.NewEncoder(w).Encode([]models.User{})
			return
		}

		log.Error("Error fetching users", sl.Err(err))
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

func (u *UsersController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.GetUsersHandler"
	log := u.log.With(
		"op", op,
	)

	id_s, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("id is required")
		http.Error(w, "Error: id is required", http.StatusBadRequest)
		return
	}

	id_uuid, err := uuid.Parse(id_s)
	if err != nil {
		log.Error("id must be UUID", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.service.GetUserById(r.Context(), id_uuid)
	if err != nil {
		if errors.Is(err, serviceerrors.ErrNotFound) {
			log.Warn("No one user with current id", sl.Err(err))
			http.Error(w, "No one user with current id", http.StatusNotFound)
			return
		}

		log.Error("Error fetching user by id", sl.Err(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

func (u *UsersController) InsertUserHandler(w http.ResponseWriter, r *http.Request) {}
func (u *UsersController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {}
func (u *UsersController) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.DeleteUserHandler"
	log := u.log.With(
		"op", op,
	)

	id_s, ok := mux.Vars(r)["id"]
	if !ok {
		log.Error("id is required")
		http.Error(w, "Error: id is required", http.StatusBadRequest)
		return
	}

	id_uuid, err := uuid.Parse(id_s)
	if err != nil {
		log.Error("id must be UUID", sl.Err(err))
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.service.DeleteUser(r.Context(), id_uuid)
	if err != nil {
		if errors.Is(err, serviceerrors.ErrNotFound) {
			log.Error("No one user with current id", sl.Err(err))
			http.Error(w, "No one user with current id", http.StatusNotFound)
			return
		}

		log.Error("Error deleting user", sl.Err(err))
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}
