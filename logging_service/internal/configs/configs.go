package configs

import "github.com/ilyakaznacheev/cleanenv"

type Configs struct {
	Server struct {
		Addr string `yaml:"addr"`
	} `yaml:"server"`

	RabbitMQ struct {
		Addr string `yaml:"addr"`
	} `yaml:"rabbitmq"`

	Postgres struct {
		Username string `yaml:"username"`
		DBName   string `yaml:"db_name"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
	} `yaml:"postgres"`
}

func GetConfigs() (*Configs, error) {
	configs := &Configs{}
	return configs, cleanenv.ReadConfig("config.yml", configs)
}
