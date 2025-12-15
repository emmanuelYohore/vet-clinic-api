package user

import (
	"net/http"
	"strconv"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/database/dbmodel"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/chi/v5"
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

	userRole := req.Role
	if userRole == "" {
		userRole = "user"
	}

	user := &dbmodel.User{
		Email:    req.Email,
		Password: req.Password,
		Role:     userRole,
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

func (config *UserConfig) GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := config.UserRepository.FindAll()
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "unable to fetch users",
		})
		return
	}
	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, users)
}

func (config *UserConfig) GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid user ID",
		})
		return
	}
	user, err := config.UserRepository.FindById(uint(userID))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "user not found",
		})
		return
	}
	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, user)
}

func (config *UserConfig) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid user ID",
		})
		return
	}
	existingUser, err := config.UserRepository.FindById(uint(userID))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "user not found",
		})
		return
	}
	if err := config.UserRepository.Delete(uint(userID), existingUser); err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to delete user",
		})
		return
	}
}

func (config *UserConfig) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userID, err := strconv.Atoi(idParam)
	if err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid user ID",
		})
		return
	}
	req := &models.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}
	existingUser, err := config.UserRepository.FindById(uint(userID))
	if err != nil {
		render.Status(r, http.StatusNotFound)
		render.JSON(w, r, map[string]string{
			"error": "user not found",
		})
		return
	}
	existingUser.Email = req.Email
	existingUser.Password = req.Password
	updatedUser, err := config.UserRepository.Update(existingUser)
	if err != nil {
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, map[string]string{
			"error": "failed to update user",
		})
		return
	}
	render.Status(r, http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	render.JSON(w, r, updatedUser)
}
