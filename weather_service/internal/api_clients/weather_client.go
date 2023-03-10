package apiclients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Levap123/playstar-test/weather_service/internal/entities"
)

type WeatherClient struct {
	Key        string
	httpClient *http.Client
}

func NewWeatherClient(key string, httpClient *http.Client) *WeatherClient {
	return &WeatherClient{
		Key:        key,
		httpClient: httpClient,
	}
}

func (wc *WeatherClient) GetWeather(ctx context.Context, city string) (entities.Weather, error) {
	URL := fmt.Sprintf("https://api.tomorrow.io/v4/weather/realtime?location=%s&apikey=%s", city, wc.Key)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, URL, nil)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("weather client - new request - %w", err)
	}

	resp, err := wc.httpClient.Do(req)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("weather client - do request - %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return entities.Weather{}, fmt.Errorf("weather client - do request - %w", ErrWeather)
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return entities.Weather{}, fmt.Errorf("weather client - read response - %w", err)
	}

	var weatherResp entities.WeatherApiResponse
	if err := json.Unmarshal(respBytes, &weatherResp); err != nil {
		return entities.Weather{}, fmt.Errorf("weather client - unmarshal response - %w", err)
	}

	return entities.FromRespToWeather(weatherResp), nil
}
