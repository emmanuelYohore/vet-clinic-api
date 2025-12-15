package treatment

import (
	"net/http"
	"strconv"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type TreatmentConfig struct {
	*config.Config
}

func New(configuration *config.Config) *TreatmentConfig {
	return &TreatmentConfig{configuration}
}

// CreateTreatmentHandler doc
// @Summary Create a new treatment
// @Tags treatments
// @Accept json
// @Produce json
// @Param treatment body models.TreatmentRequest true "Treatment payload"
// @Success 201 {object} dbmodel.Treatment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /treatments [post]
func (config *TreatmentConfig) CreateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.TreatmentRequest{}

	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	treatment := &dbmodel.Treatment{
		Name: req.Name,
	}

	savedTreatment, err := config.TreatmentRepository.Create(treatment)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to save treatment",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, savedTreatment)
}

// GetAllTreatmentsHandler doc
// @Summary Get all treatments
// @Tags treatments
// @Produce json
// @Success 200 {array} dbmodel.Treatment
// @Failure 500 {object} map[string]string
// @Router /treatments [get]
func (config *TreatmentConfig) GetAllTreatmentsHandler(w http.ResponseWriter, r *http.Request) {
	treatments, err := config.TreatmentRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch treatments",
		})
		return
	}

	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, treatments)
}

// GetTreatmentByIDHandler doc
// @Summary Get a treatment by ID
// @Tags treatments
// @Produce json
// @Param id path int true "Treatment ID"
// @Success 200 {object} dbmodel.Treatment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /treatments/{id} [get]
func (config *TreatmentConfig) GetTreatmentByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid treatment ID",
		})
		return
	}

	treatment, err := config.TreatmentRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "cat not found",
		})
		return
	}
	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, treatment)
}

// UpdateTreatmentHandler doc
// @Summary Update a treatment
// @Tags treatments
// @Accept json
// @Produce json
// @Param id path int true "Treatment ID"
// @Param treatment body models.TreatmentRequest true "Treatment payload"
// @Success 200 {object} dbmodel.Treatment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /treatments/{id} [put]
func (config *TreatmentConfig) UpdateTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid treatment ID",
		})
		return
	}

	req := &models.TreatmentRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	existing, err := config.TreatmentRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "treatment not found",
		})
		return
	}

	existing.Name = req.Name

	updatedTreatment, err := config.TreatmentRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to update treatment",
		})
		return
	}
	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, updatedTreatment)
}

// DeleteTreatmentHandler doc
// @Summary Delete a treatment
// @Tags treatments
// @Param id path int true "Treatment ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /treatments/{id} [delete]
func (config *TreatmentConfig) DeleteTreatmentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid treatment ID",
		})
		return
	}

	treatment, err := config.TreatmentRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "treatment not found",
		})
		return
	}

	if err := config.TreatmentRepository.Delete(uint(id64), treatment); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to delete treatment",
		})
		return
	}
	
	render.Status(r, http.StatusNoContent)
}

// GetTreatmentByVisitHandler doc
// @Summary Get treatments by Visit ID
// @Tags treatments
// @Produce json
// @Param visit_id path int true "Visit ID"
// @Success 200 {array} dbmodel.Treatment
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /visits/{id}/treatments [get]
func (config *TreatmentConfig) GetTreatmentByVisitHandler(w http.ResponseWriter, r *http.Request) {
	visitIDParam := chi.URLParam(r, "visit_id")
	visitID64, err := strconv.ParseUint(visitIDParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid visit ID",
		})
		return
	}
	treatments, err := config.TreatmentRepository.FindByVisitID(uint(visitID64))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch treatments for the visit",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")

	render.JSON(w, r, treatments)
}
