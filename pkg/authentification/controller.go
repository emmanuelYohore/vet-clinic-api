package authentification

import (
	"encoding/json"
	"net/http"

	"github.com/emmanuelYohore/vet-clinic-api/config"
	"github.com/emmanuelYohore/vet-clinic-api/pkg/models"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AuthConfig {
	return &AuthConfig{configuration}
}

func (c *AuthConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	payload := &models.UserRequest{}

	if err := render.Bind(r, payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	user, err := c.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	userRole := user.Role
	if userRole == "" {
		userRole = "user"
	}

	token, err := GenerateToken("your_secret_key", payload.Email, userRole)
	refreshToken, _ := GenerateRefreshToken(payload.Email, userRole)

	user.RefreshToken = refreshToken
	c.UserRepository.Update(user)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token, "refresh_token": refreshToken})
}

func (c *AuthConfig) RefreshToken(w http.ResponseWriter, r *http.Request) {
	payload := &models.UserRequest{}
	if err := render.Bind(r, payload); err != nil {
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, map[string]string{
			"error": "invalid request payload",
		})
		return
	}

	user, err := c.UserRepository.GetUserByEmail(payload.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	userRole := user.Role
	if userRole == "" {
		userRole = "user"
	}

	newToken, err := GenerateToken("your_secret_key", payload.Email, userRole)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": newToken})
}
