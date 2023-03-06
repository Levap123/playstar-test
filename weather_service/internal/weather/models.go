package weather

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}
