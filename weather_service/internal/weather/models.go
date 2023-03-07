package weather

import "github.com/Levap123/playstar-test/weather_service/proto"

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}

type Coordinates struct {
	Latitude  float32
	Longitude float32
}

func FromProtoRequestToCoordinates(req *proto.GetWeatherRequest) Coordinates {
	return Coordinates{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
