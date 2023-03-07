package apiclients

import (
	"context"

	"github.com/Levap123/playstar-test/weather_service/proto"
)

type CityClient struct {
	cl proto.CityServiceClient
}

func (cc *CityClient) GetCityByCoordinates(ctx context.Context, latitude, longitude float32) (string, error) {
	req := &proto.GetCityRequest{
		Latitude:  latitude,
		Longitude: longitude,
	}

	resp, err := cc.cl.GetCity(ctx, req)
	if err != nil {
		return "", err
	}

	return resp.City, nil
}
