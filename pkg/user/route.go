package user

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	userConfig := New(configuration)

	router := chi.NewRouter()

	router.Post("/", userConfig.CreateUserHandler)

	return router
}
