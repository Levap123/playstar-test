package entities

type Weather struct {
	Temperature              float32 `json:"temperature"`
	TemperatureApparent      float32 `json:"temperatureApparent"`
	Humidity                 uint64  `json:"humidity"`
	PrecipitationProbability uint64  `json:"precipitationProbability"`
}

type WeatherApiResponse struct {
	Data struct {
		Values struct {
			Temperature              float32 `json:"temperature,omitempty"`
			TemperatureApparent      float32 `json:"temperatureApparent,omitempty"`
			Humidity                 uint64  `json:"humidity,omitempty"`
			PrecipitationProbability uint64  `json:"precipitationProbability,omitempty"`
		}
	}
}

func FromRespToWeather(resp WeatherApiResponse) Weather {
	return Weather{
		Temperature:              resp.Data.Values.Temperature,
		TemperatureApparent:      resp.Data.Values.TemperatureApparent,
		Humidity:                 resp.Data.Values.Humidity,
		PrecipitationProbability: resp.Data.Values.PrecipitationProbability,
	}
}
