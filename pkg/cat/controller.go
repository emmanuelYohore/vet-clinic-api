package cat

import (
	"net/http"
	"strconv"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type CatConfig struct {
	*config.Config
}

func New(configuration *config.Config) *CatConfig {
	return &CatConfig{configuration}
}

// CreateCatHandler godoc
// @Summary Create a cat
// @Tags cats
// @Accept json
// @Produce json
// @Param cat body models.CatRequest true "Cat payload"
// @Success 201 {object} dbmodel.Cat
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats [post]
func (config *CatConfig) CreateCatHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.CatRequest{}

	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	cat := &dbmodel.Cat{
		Name:   req.Name,
		Age:    req.Age,
		Breed:  req.Breed,
		Weigth: req.Weigth,
	}

	savedCat, err := config.CatRepository.Create(cat)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to save cat",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, savedCat)
}

// GetAllCatsHandler godoc
// @Summary List cats
// @Tags cats
// @Produce json
// @Success 200 {array} dbmodel.Cat
// @Failure 500 {object} map[string]string
// @Router /cats [get]
func (config *CatConfig) GetAllCatsHandler(w http.ResponseWriter, r *http.Request) {
	cats, err := config.CatRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to fetch cats",
		})
		return
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, cats)
}

// GetCatByIDHandler godoc
// @Summary Get a cat by ID
// @Tags cats
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} dbmodel.Cat
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /cats/{id} [get]
func (config *CatConfig) GetCatByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid cat ID",
		})
		return
	}

	cat, err := config.CatRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "cat not found",
		})
		return
	}

	render.JSON(w, r, cat)
}

// UpdateCatHandler godoc
// @Summary Update a cat
// @Tags cats
// @Accept json
// @Produce json
// @Param id path int true "Cat ID"
// @Param cat body models.CatRequest true "Cat payload"
// @Success 200 {object} dbmodel.Cat
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id} [put]
func (config *CatConfig) UpdateCatHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid cat ID",
		})
		return
	}

	req := &models.CatRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	existing, err := config.CatRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "cat not found",
		})
		return
	}

	existing.Name = req.Name
	existing.Age = req.Age
	existing.Breed = req.Breed
	existing.Weigth = req.Weigth

	updatedCat, err := config.CatRepository.Update(existing)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to update cat",
		})
		return
	}

	render.JSON(w, r, updatedCat)
}

// DeleteCatHandler godoc
// @Summary Delete a cat
// @Tags cats
// @Param id path int true "Cat ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id} [delete]
func (config *CatConfig) DeleteCatHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id64, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid cat ID",
		})
		return
	}

	cat, err := config.CatRepository.FindById(uint(id64))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "cat not found",
		})
		return
	}

	if err := config.CatRepository.Delete(uint(id64), cat); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to delete cat",
		})
		return
	}

	render.Status(r, http.StatusNoContent)
	render.JSON(w, r, nil)
}

// GetCatHistoryHandler godoc
// @Summary Get a cat history (visits)
// @Tags cats
// @Produce json
// @Param id path int true "Cat ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /cats/{id}/history [get]
func (c *CatConfig) GetCatHistoryHandler(w http.ResponseWriter, r *http.Request) {
    idParam := chi.URLParam(r, "id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        render.Status(r, http.StatusBadRequest)
        render.JSON(w, r, map[string]string{"error": "invalid cat ID"})
        return
    }

    cat, err := c.Config.CatRepository.FindById(uint(id))
    if err != nil {
        render.Status(r, http.StatusNotFound)
        render.JSON(w, r, map[string]string{"error": "cat not found"})
        return
    }

    visits, err := c.Config.VisitRepository.FindByCatID(uint(id))
    if err != nil {
        render.Status(r, http.StatusInternalServerError)
        render.JSON(w, r, map[string]string{"error": "could not load visits"})
        return
    }

    history := map[string]interface{}{
        "cat":    cat,
        "visits": visits,
    }

    render.Status(r, http.StatusOK)
    render.JSON(w, r, history)
}