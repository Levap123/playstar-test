package apiclients

import (
	"context"

	"github.com/Levap123/playstar-test/api_gateway/internal/dto"
	"github.com/Levap123/playstar-test/api_gateway/proto"
	"google.golang.org/grpc"
)

type WeatherClient struct {
	cl proto.WeatherServiceClient
}

func NewWeatherClient(conn *grpc.ClientConn) *WeatherClient {
	cl := proto.NewWeatherServiceClient(conn)
	return &WeatherClient{
		cl: cl,
	}
}

func (wc *WeatherClient) GetWeather(ctx context.Context, dtoWeather dto.GetWeatherDTO) (dto.Weather, error) {
	req := dto.FromDtoToReq(dtoWeather)

	resp, err := wc.cl.GetWeather(ctx, req)
	if err != nil {
		return dto.Weather{}, err
	}

	return dto.FromRespToDTO(resp), nil
}
