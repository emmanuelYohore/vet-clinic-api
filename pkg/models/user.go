package models

import (
	"errors"
	"net/http"
)

type UserRequest struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	
}

func (u *UserRequest) Bind(r *http.Request) error {
	if u.Email == "" {
		return errors.New("le champ email ne doit pas être vide")
	}
	if u.Password == "" {

		return errors.New("le champ password ne doit pas être vide")
	}

	
	return nil

}

type UserResponse struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	RefreshToken string `json:"refresh_token"`
}
