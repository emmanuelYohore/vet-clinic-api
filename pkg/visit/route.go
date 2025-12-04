package visit

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	visitConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", visitConfig.CreateVisitHandler)
	router.Get("/", visitConfig.GetAllVisitsHandler)
	router.Get("/{id}", visitConfig.GetVisitByIDHandler)
	router.Put("/{id}", visitConfig.UpdateVisitHandler)
	router.Delete("/{id}", visitConfig.DeleteVisitHandler)
	router.Get("/cats/{id}/visits", visitConfig.GetVisitsByCatHandler)
	
	

	

	return router
}
