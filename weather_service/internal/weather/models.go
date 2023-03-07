package weather

import "github.com/Levap123/playstar-test/weather_service/proto"

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}

