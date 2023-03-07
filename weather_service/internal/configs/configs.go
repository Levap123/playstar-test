package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	WeatherApi struct {
		Key string `yaml:"key"`
	} `yaml:"weather_api"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
