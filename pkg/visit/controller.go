package visit

import (
	"net/http"
	"strconv"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type VisitConfig struct {
	*config.Config
}

func New(configuration *config.Config) *VisitConfig {
	return &VisitConfig{configuration}
}

// CreateVisitHandler doc
// @Summary Create a new visit
// @Tags visits
// @Accept json
// @Produce json
// @Param visit body models.VisitRequest true "Visit payload"
// @Success 201 {object} dbmodel.Visit
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /visits [post]
func (config *VisitConfig) CreateVisitHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.VisitRequest{}

	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	visit := &dbmodel.Visit{
		Motif:       req.Motif,
		Date:        req.Date,
		Veterinaire: req.Veterinaire,
	}

	savedVisit, err := config.VisitRepository.Create(visit)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to save visit",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, savedVisit)
}

// GetAllVisitsHandler doc
// @Summary Get all visits
// @Tags visits
// @Produce json
// @Success 200 {array} dbmodel.Visit
// @Failure 500 {object} map[string]string
// @Router /visits [get]
func (config *VisitConfig) GetAllVisitsHandler(w http.ResponseWriter, r *http.Request) {
	visits, err := config.TreatmentRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch visits",
		})
		return
	}

	render.JSON(w, r, visits)
}

// GetVisitByIDHandler doc
// @Summary Get a visit by ID
// @Tags visits
// @Produce json
// @Param id path int true "Visit ID"
// @Success 200 {object} dbmodel.Visit
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /visits/{id} [get]
func (config *VisitConfig) GetVisitByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid visit ID",
		})
		return
	}

	visit, err := config.VisitRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "cat not found",
		})
		return
	}

	render.JSON(w, r, visit)
}

// UpdateVisitHandler doc
// @Summary Update a visit
// @Tags visits
// @Accept json
// @Produce json
// @Param id path int true "Visit ID"
// @Param visit body models.VisitRequest true "Visit payload"
// @Success 200 {object} dbmodel.Visit
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /visits/{id} [put]
func (config *VisitConfig) UpdateVisitHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid visit ID",
		})
		return
	}

	req := &models.VisitRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	existing, err := config.VisitRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "visit not found",
		})
		return
	}

	existing.Motif = req.Motif
	existing.Date = req.Date
	existing.Veterinaire = req.Veterinaire

	updatedVisit, err := config.VisitRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to update visit",
		})
		return
	}

	render.JSON(w, r, updatedVisit)
}

// DeleteVisitHandler doc
// @Summary Delete a visit
// @Tags visits
// @Param id path int true "Visit ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /visits/{id} [delete]
func (config *VisitConfig) DeleteVisitHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid visit ID",
		})
		return
	}

	visit, err := config.VisitRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "visit not found",
		})
		return
	}

	if err := config.VisitRepository.Delete(uint(id64), visit); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to delete visit",
		})
		return
	}

	render.Status(r, http.StatusNoContent)
}

// GetVisitsByCatHandler doc
// @Summary Get visits by Cat ID
// @Tags visits
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {array} dbmodel.Visit
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id}/visits [get]
func (config *VisitConfig) GetVisitsByCatHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)	
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid cat ID",
		})
		return
	}
	visits, err := config.VisitRepository.FindByCatID(uint(id64))
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch visits for cat",
		})
		return
	}
	render.JSON(w, r, visits)
}

// FilterByMotifOrVeterinaireHandler doc
// @Summary Filter visits by motif or veterinaire
// @Tags visits
// @Produce json
// @Param motif query string false "Motif"
// @Param veterinaire query string false "Veterinaire"
// @Success 200 {array} dbmodel.Visit
// @Failure 500 {object} map[string]string
// @Router /visits/filter [get]
func (config *VisitConfig) FilterByMotifOrVeterinaireHandler(w http.ResponseWriter, r *http.Request) {
	motif := r.URL.Query().Get("motif")
	veterinaire := r.URL.Query().Get("veterinaire")
	visits, err := config.VisitRepository.FilterByMotifOrVeterinaire(motif, veterinaire)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to filter visits",
		})
		return
	}
	render.JSON(w, r, visits)
}
