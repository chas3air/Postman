package userscontroller

import (
	"api/internal/domain/interfaces/service"
	"api/internal/domain/models"
	serviceerrors "api/internal/service"
	"api/pkg/lib/logger/sl"
	"encoding/json"
	"errors"
	"io"
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

func (u *UsersController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.GetUserByIdHandler"
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

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

func (u *UsersController) InsertUserHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.InsertUserHandler"
	log := u.log.With(
		"op", op,
	)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Cannot read request body", sl.Err(err))
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Error("Error parse request body to object", sl.Err(err))
		http.Error(w, "Error parse request body to object", http.StatusBadRequest)
		return
	}

	inserted_user, err := u.service.InsertUser(r.Context(), user)
	if err != nil {
		if errors.Is(err, serviceerrors.ErrAlreadyExists) {
			log.Error("User already exists", sl.Err(err))
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(inserted_user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

func (u *UsersController) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	const op = "controller.UpdateUserHandler"
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

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("Cannot read request body", sl.Err(err))
		http.Error(w, "Cannot read request body", http.StatusBadRequest)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		log.Error("Error parse request body to object", sl.Err(err))
		http.Error(w, "Error parse request body to object", http.StatusBadRequest)
		return
	}

	updated_user, err := u.service.UpdateUser(r.Context(), id_uuid, user)
	if err != nil {
		if errors.Is(err, serviceerrors.ErrNotFound) {
			log.Warn("No one user with current id", sl.Err(err))
			http.Error(w, "No one user with current id", http.StatusNotFound)
			return
		}

		log.Error("Error updating user", sl.Err(err))
		http.Error(w, "Error updating user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(updated_user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}

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
			log.Warn("No one user with current id", sl.Err(err))
			http.Error(w, "No one user with current id", http.StatusNotFound)
			return
		}

		log.Error("Error deleting user", sl.Err(err))
		http.Error(w, "Error deleting user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Error("Error write users to response body", sl.Err(err))
		http.Error(w, "Error write users to response body", http.StatusInternalServerError)
		return
	}
}
