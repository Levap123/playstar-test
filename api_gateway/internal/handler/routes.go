package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewMux()

	r.Get("api/v1/weather", h.getWeatherByCoordinates)

	return r
}
