package handler

import (
	"context"
	"errors"
	"time"

	apiclients "github.com/Levap123/playstar-test/weather_service/internal/api_clients"
	"github.com/Levap123/playstar-test/weather_service/internal/logs"
	"github.com/Levap123/playstar-test/weather_service/internal/validator"
	"github.com/Levap123/playstar-test/weather_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WeatherHandler struct {
	proto.UnimplementedWeatherServiceServer
	cc     *apiclients.CityClient
	wc     *apiclients.WeatherClient
	v      validator.Validator
	logger *logs.Logger
}

func NewWeatherHandler(cc *apiclients.CityClient, wc *apiclients.WeatherClient, v validator.Validator, logger *logs.Logger) *WeatherHandler {
	return &WeatherHandler{
		cc:     cc,
		wc:     wc,
		v:      v,
		logger: logger,
	}
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
	h.logger.Info().Msg("get weather handler")

	cityCtx, cityCancel := context.WithTimeout(ctx, time.Second)
	defer cityCancel()

	if err := h.v.ValidateCoordinates(req.Latitude, req.Longitude); err != nil {
		h.logger.Err(err).Msg("coordinates incorrect")
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	city, err := h.cc.GetCityByCoordinates(cityCtx, req.Latitude, req.Longitude)
	if err != nil {
		h.logger.Err(err).Msg("something went wrong in getting city by coordinates")
		return nil, err
	}

	weatherCtx, weatherCancel := context.WithTimeout(ctx, time.Second)
	defer weatherCancel()

	weather, err := h.wc.GetWeather(weatherCtx, city)
	if err != nil {
		h.logger.Err(err).Msg("something went wrong in getting weather by city")

		if errors.Is(err, apiclients.ErrWeather) {
			return nil, status.Error(codes.Aborted, apiclients.ErrWeather.Error())
		}
		return nil, err
	}

	h.logger.Info().Msg("success in sending response")

	return &proto.GetWeatherResponse{
		Temperature:              weather.Temperature,
		TemperatureApparent:      weather.Temperature,
		Humidity:                 weather.Humidity,
		PrecipitationProbability: weather.PrecipitationProbability,
	}, nil
}
