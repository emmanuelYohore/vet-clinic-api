package treatment

import (
	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/go-chi/chi/v5"
)

func Routes(configuration *config.Config) *chi.Mux {
	treatmentConfig := New(configuration)
	router := chi.NewRouter()

	router.Post("/", treatmentConfig.CreateTreatmentHandler)
	router.Get("/", treatmentConfig.GetAllTreatmentsHandler)
	router.Get("/{id}", treatmentConfig.GetTreatmentByIDHandler)

	router.Get("/visits/{id}/treatments", treatmentConfig.GetTreatmentByVisitHandler)
	router.Put("/{id}", treatmentConfig.UpdateTreatmentHandler)
	router.Delete("/{id}", treatmentConfig.DeleteTreatmentHandler)
	


	return router
}
