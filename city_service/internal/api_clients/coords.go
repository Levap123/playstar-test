package apiclients

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Levap123/playstar-test/city_service/internal/city"
)

type CoordinatesClient struct {
	httpClient *http.Client
}

func NewCoordinatesClient(cl *http.Client) *CoordinatesClient {
	return &CoordinatesClient{
		httpClient: cl,
	}
}

func (cc *CoordinatesClient) GetCityFromCoordiantes(latitude, longitude float32) (city.City, error) {
	URL := fmt.Sprintf(`https://api.bigdatacloud.net/data/reverse-geocode-client?latitude=%f&longitude=%f&localityLanguage=en`,
		latitude, longitude)
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return city.City{}, fmt.Errorf("coordinates client - get city - new request - %w", err)
	}

	resp, err := cc.httpClient.Do(req)
	if err != nil {
		return city.City{}, fmt.Errorf("coordinates client - get city - do request - %w", err)
	}
	defer resp.Body.Close()

	reqBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return city.City{}, fmt.Errorf("coordinates client - get city - read request - %w", err)
	}

	var cityResp city.City
	if err := json.Unmarshal(reqBytes, &cityResp); err != nil {
		return city.City{}, fmt.Errorf("coordinates client - get city - unamrshal - %w", err)
	}

	return cityResp, nil
}
