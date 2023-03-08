package handler

import apiclients "github.com/Levap123/playstar-test/api_gateway/internal/api_clients"

type Handler struct {
	wc *apiclients.WeatherClient
}

func NewHandler(wc *apiclients.WeatherClient) *Handler {
	return &Handler{
		wc: wc,
	}
}
