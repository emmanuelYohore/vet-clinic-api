package models

import (
	"errors"
	"net/http"
)

type CatRequest struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Breed  string `json:"breed"`
	Weigth int    `json:"weigth"`
}

func (c *CatRequest) Bind(r *http.Request) error{
	if c.Name == "" {
		return errors.New("le champ name ne doit pas être vide")
	}
	if c.Age < 0 {
		return errors.New("age doit être supérieur ou égale à 0 ")
	}
	if c.Breed == "" {
		return errors.New("le champ breed ne doit pas être vide")
	}
	if c.Age < 0 {
		return errors.New("weigth doit être supérieur ou égale à 0 ")
	}
	return nil

}

type CatResponse struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Breed  string `json:"breed"`
	Weigth int    `json:"weigth"`
}
