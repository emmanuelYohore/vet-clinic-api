package authentification

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	authConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", authConfig.LoginHandler)
	router.Post("/refresh", authConfig.RefreshToken)

	return router
}
