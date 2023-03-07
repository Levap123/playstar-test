package weather

import (
	"context"
	"errors"
	"time"

	apiclients "github.com/Levap123/playstar-test/weather_service/internal/api_clients"
	"github.com/Levap123/playstar-test/weather_service/internal/validator"
	"github.com/Levap123/playstar-test/weather_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WeatherHandler struct {
	proto.UnimplementedWeatherServiceServer
	cc *apiclients.CityClient
	wc *apiclients.WeatherClient
	v  validator.Validator
}

func NewWeatherHandler(cc *apiclients.CityClient, wc *apiclients.WeatherClient, v validator.Validator) *WeatherHandler {
	return &WeatherHandler{
		cc: cc,
		wc: wc,
		v:  v,
	}
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
	cityCtx, cityCancel := context.WithTimeout(ctx, time.Second)
	defer cityCancel()

	if err := h.v.ValidateCoordinates(req.Latitude, req.Longitude); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	city, err := h.cc.GetCityByCoordinates(cityCtx, req.Latitude, req.Longitude)
	if err != nil {
		return nil, err
	}

	weatherCtx, weatherCancel := context.WithTimeout(ctx, time.Second)
	defer weatherCancel()

	weather, err := h.wc.GetWeather(weatherCtx, city)
	if err != nil {
		if errors.Is(err, apiclients.ErrWeather) {
			return nil, status.Error(codes.Aborted, apiclients.ErrWeather.Error())
		}
		return nil, err
	}

	return &proto.GetWeatherResponse{
		Temperature:              weather.Temperature,
		TemperatureApparent:      weather.Temperature,
		Humidity:                 weather.Humidity,
		PrecipitationProbability: weather.PrecipitationProbability,
	}, nil
}
