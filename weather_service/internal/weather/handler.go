package weather

import (
	"context"
	"time"

	apiclients "github.com/Levap123/playstar-test/weather_service/internal/api_clients"
	"github.com/Levap123/playstar-test/weather_service/proto"
)

type WeatherHandler struct {
	proto.UnimplementedWeatherServiceServer
	cc *apiclients.CityClient
	wc *apiclients.WeatherClient
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
	cityCtx, cityCancel := context.WithTimeout(ctx, time.Second)
	defer cityCancel()

	city, err := h.cc.GetCityByCoordinates(cityCtx, req.Latitude, req.Longitude)
	if err != nil {
		return nil, err
	}

	weatherCtx, weatherCancel := context.WithTimeout(ctx, time.Second)
	defer weatherCancel()

	weather, err := h.wc.GetWeather(weatherCtx, city)
	if err != nil {
		return nil, err
	}

	return &proto.GetWeatherResponse{
		Temperature:              weather.Temperature,
		TemperatureApparent:      weather.Temperature,
		Humidity:                 weather.Humidity,
		PrecipitationProbability: weather.PrecipitationProbability,
	}, nil
}
