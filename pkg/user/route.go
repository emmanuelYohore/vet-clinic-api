package user

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	userConfig := New(configuration)

	router := chi.NewRouter()

	router.Post("/", userConfig.CreateUserHandler)
	router.Get("/", userConfig.GetAllUsersHandler)
	router.Get("/{id}", userConfig.GetUserByIDHandler)
	router.Put("/{id}", userConfig.UpdateUserHandler)
	router.Delete("/{id}", userConfig.DeleteUserHandler)

	return router
}
