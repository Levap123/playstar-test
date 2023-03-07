package apiclients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CoordinatesClient struct {
	httpClient *http.Client
}

type City struct {
	City string `json:"city,omitempty"`
}

func NewCoordinatesClient(cl *http.Client) *CoordinatesClient {
	return &CoordinatesClient{
		httpClient: cl,
	}
}

func (cc *CoordinatesClient) GetCityFromCoordiantes(ctx context.Context, latitude, longitude float32) (City, error) {
	URL := fmt.Sprintf(`https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%f&longitude=%f&localityLanguage=en`,
		latitude, longitude)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return City{}, fmt.Errorf("coordinates client - get city - new request - %w", err)
	}

	resp, err := cc.httpClient.Do(req)
	if err != nil {
		return City{}, fmt.Errorf("coordinates client - get city - do request - %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return City{}, fmt.Errorf("coordinates client - get city - do request - %w", ErrBadRequest)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return City{}, fmt.Errorf("coordinates client - get city - do request - %w", ErrInternal)
	}

	reqBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return City{}, fmt.Errorf("coordinates client - get city - read request - %w", err)
	}

	var cityResp City
	if err := json.Unmarshal(reqBytes, &cityResp); err != nil {
		return City{}, fmt.Errorf("coordinates client - get city - unamrshal - %w", err)
	}

	return cityResp, nil
}
