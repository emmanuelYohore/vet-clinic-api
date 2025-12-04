package models

import (
	"errors"
	"net/http"
)

type TreatmentRequest struct {
	Name string `json:"name"`
}

func (t *TreatmentRequest) Bind(r *http.Request) error {
	if t.Name == "" {
		return errors.New("le champ name ne doit pas être vide")
	}
	return nil
}

type TreatmentResponse struct {
	Name string `json:"name"`
}
