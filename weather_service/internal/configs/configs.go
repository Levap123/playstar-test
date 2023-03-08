package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	WeatherApi struct {
		Key string `yaml:"key"`
	} `yaml:"weather_api"`

	Client struct {
		TimeoutSeconds int `yaml:"timeout"`
	} `yaml:"client"`

	CityService struct {
		Addr string `yaml:"addr"`
	} `yaml:"city_service"`

	RabbitMQ struct {
		Addr string `yaml:"addr"`
	} `yaml:"rabbitmq"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
