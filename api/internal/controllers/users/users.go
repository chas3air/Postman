package userscontroller

import (
	"api/internal/domain/interfaces/service"
	"log/slog"
	"net/http"
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

func (u *UsersController) GetUsersHandler(w http.ResponseWriter, r *http.Request)    {}
func (u *UsersController) GetUserByIdHandler(w http.ResponseWriter, r *http.Request) {}
func (u *UsersController) InsertUserHandler(w http.ResponseWriter, r *http.Request)  {}
func (u *UsersController) UpdateUserHandler(w http.ResponseWriter, r *http.Request)  {}
func (u *UsersController) DeleteUserHandler(w http.ResponseWriter, r *http.Request)  {}
