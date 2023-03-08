package apiclients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Levap123/playstar-test/city_service/internal/entities"
)

type CoordinatesClient struct {
	httpClient *http.Client
}

func NewCoordinatesClient(cl *http.Client) *CoordinatesClient {
	return &CoordinatesClient{
		httpClient: cl,
	}
}

func (cc *CoordinatesClient) GetCityFromCoordiantes(ctx context.Context, latitude, longitude float32) (entities.City, error) {
	URL := fmt.Sprintf(`https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%f&longitude=%f&localityLanguage=en`,
		latitude, longitude)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return entities.City{}, fmt.Errorf("coordinates client - get city - new request - %w", err)
	}

	resp, err := cc.httpClient.Do(req)
	if err != nil {
		return entities.City{}, fmt.Errorf("coordinates client - get city - do request - %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusBadRequest {
		return entities.City{}, fmt.Errorf("coordinates client - get city - do request - %w", ErrBadRequest)
	}

	if resp.StatusCode == http.StatusInternalServerError {
		return entities.City{}, fmt.Errorf("coordinates client - get city - do request - %w", ErrInternal)
	}

	reqBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.City{}, fmt.Errorf("coordinates client - get city - read request - %w", err)
	}

	var cityResp entities.City
	if err := json.Unmarshal(reqBytes, &cityResp); err != nil {
		return entities.City{}, fmt.Errorf("coordinates client - get city - unamrshal - %w", err)
	}

	return cityResp, nil
}
