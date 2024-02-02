package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type Service interface {
	GetCatFact(context.Context) (*CatFact, error)
}

type CatFactService struct {
	url string
}

func NewCatFactService(url string) Service {
	return &CatFactService{
		url: url,
	}
}

func (s *CatFactService) GetCatFact(context.Context) (*CatFact, error) {

	res, err := http.Get(s.url)

	if err != nil {
		return &CatFact{}, err
	}

	var data = new(CatFact)

	if err := json.NewDecoder(res.Body).Decode(data); err != nil {
		return &CatFact{}, err
	}

	return data, nil

}
