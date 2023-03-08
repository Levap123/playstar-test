package handler

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/Levap123/playstar-test/api_gateway/internal/dto"
)

func (h *Handler) getWeatherByCoordinates(w http.ResponseWriter, r *http.Request) {
	var dto dto.GetWeatherDTO

	reqBytes, err := io.ReadAll(r.Body)
	if err != nil {
		errBytes := marshal(map[string]string{"error": err.Error()})
		sendJSON(w, errBytes, http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(reqBytes, &dto); err != nil {
		errBytes := marshal(map[string]string{"error": err.Error()})
		sendJSON(w, errBytes, http.StatusBadRequest)
		return

	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	resp, err := h.wc.GetWeather(ctx, dto)
	if err != nil {
		errBytes := marshal(map[string]string{"error": err.Error()})
		sendJSON(w, errBytes, http.StatusBadRequest)
		return

	}

	respBytes := marshal(resp)

	sendJSON(w, respBytes, http.StatusOK)
}

func sendJSON(w http.ResponseWriter, responseBytes []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(responseBytes)
}

func marshal(response any) []byte {
	requestBytes, err := json.Marshal(response)

	if err != nil {
		return nil
	}

	return requestBytes
}
