package app

import (
	userscontroller "api/internal/controllers/users"
	"api/internal/domain/interfaces/storage"
	usersservice "api/internal/service/users"
	"api/pkg/lib/logger/sl"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	log     *slog.Logger
	port    int
	storage storage.IUsersStorage
}

func New(log *slog.Logger, port int, storage storage.IUsersStorage) *App {
	return &App{
		log:     log,
		port:    port,
		storage: storage,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "app.Run"
	usersService := usersservice.New(a.log, a.storage)
	usersController := userscontroller.New(a.log, usersService)
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/health-check", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Ok"))
		w.WriteHeader(http.StatusOK)
	})

	r.HandleFunc("/api/v1/users", usersController.GetUsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users/{id}", usersController.GetUserByIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/users", usersController.InsertUserHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/users/{id}", usersController.UpdateUserHandler).Methods(http.MethodPut)
	r.HandleFunc("/api/v1/users/{id}", usersController.DeleteUserHandler).Methods(http.MethodDelete)

	if err := http.ListenAndServe(fmt.Sprintf(":%d", a.port), r); err != nil {
		a.log.With("op", op).Error("Error run server", sl.Err(err))
		panic(err)
	}

	return nil
}
