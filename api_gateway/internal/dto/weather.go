package dto

import "github.com/Levap123/playstar-test/api_gateway/proto"

type GetWeatherDTO struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
}

func FromDtoToReq(dto GetWeatherDTO) *proto.GetWeatherRequest {
	return &proto.GetWeatherRequest{
		Latitude:  dto.Latitude,
		Longitude: dto.Longitude,
	}
}

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}

func FromRespToDTO(req *proto.GetWeatherResponse) Weather {
	return Weather{
		Temperature:              req.Temperature,
		TemperatureApparent:      req.TemperatureApparent,
		Humidity:                 req.Humidity,
		PrecipitationProbability: req.PrecipitationProbability,
	}
}
