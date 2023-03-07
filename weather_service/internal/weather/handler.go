package weather

import (
	"context"

	"github.com/Levap123/playstar-test/weather_service/proto"
)

type WeatherHandler struct {
	proto.UnimplementedWeatherServiceServer
}

func (h *WeatherHandler) GetWeather(ctx context.Context, req *proto.GetWeatherRequest) (*proto.GetWeatherResponse, error) {
}
