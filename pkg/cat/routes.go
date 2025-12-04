package cat

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	catConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", catConfig.CreateCatHandler)
	router.Get("/", catConfig.GetAllCatsHandler)
	router.Get("/{id}", catConfig.GetCatByIDHandler)
	router.Put("/{id}", catConfig.UpdateCatHandler)
	router.Delete("/{id}", catConfig.DeleteCatHandler)
	router.Get("/{id}/history", catConfig.GetCatHistoryHandler)

	return router
}
