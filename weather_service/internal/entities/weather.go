package entities

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}

type WeatherApiResponse struct {
	Data struct {
		Temperature              float32 `json:"temperature,omitempty"`
		TemperatureApparent      float32 `json:"temperatureApparent,omitempty"`
		Humidity                 uint64  `json:"humidity,omitempty"`
		PrecipitationProbability uint64  `json:"precipitationProbability,omitempty"`
	} `json:"data,omitempty"`
}

func FromRespToWeather(resp WeatherApiResponse) Weather {
	return Weather{
		Temperature:              resp.Data.Temperature,
		TemperatureApparent:      resp.Data.TemperatureApparent,
		Humidity:                 resp.Data.Humidity,
		PrecipitationProbability: resp.Data.PrecipitationProbability,
	}
}
