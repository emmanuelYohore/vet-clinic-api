package models

import (
	"errors"
	"net/http"
	"time"
)

type VisitRequest struct {
	Date        time.Time `json:"date"`
	Motif       string    `json:"motif"`
	Veterinaire string    `json:"veterinaire"`
}

func (v *VisitRequest)  Bind(r *http.Request) error{
	if v.Motif == "" {
		return errors.New("le champ motif ne doit pas être vide")
	}

	if v.Veterinaire== "" {
		return errors.New("le champ veterinaire ne doit pas être vide")
	}

	if v.Date.IsZero() {
		return errors.New("le champ date ne doit pas être vide")
	}

	return nil
}

type VisitResponse struct {
	Date        time.Time `json:"date"`
	Motif       string    `json:"motif"`
	Veterinaire string    `json:"veterinaire"`
}
