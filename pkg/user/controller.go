package user

import (
	"net/http"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/render"
)

type UserConfig struct {
	*config.Config
}

func New(configuration *config.Config) *UserConfig {
	return &UserConfig{configuration}
}

func (config *UserConfig) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := &models.UserRequest{}

	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	user := &dbmodel.User{
		Email:    req.Email,
		Password: req.Password,
	}

	savedUser, err := config.UserRepository.Create(user)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to save user",
		})
		return
	}

	render.Status(r, http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, savedUser)
}

