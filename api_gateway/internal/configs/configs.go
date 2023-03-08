package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	WeatherService struct {
		Addr string `yaml:"addr"`
	} `yaml:"weather_service"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
