package apiclients

import "net/http"

type CoordinatesClient struct {
	httpClient *http.Client
}

func NewCoordinatesClient(cl *http.Client) *CoordinatesClient {
	return &CoordinatesClient{
		httpClient: cl,
	}
}

func (cc *CoordinatesClient) GetCityFromCoordiantes(latitude, longitude float32) {
	
}
