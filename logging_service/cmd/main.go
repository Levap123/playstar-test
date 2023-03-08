package main

import (
	"github.com/Levap123/playstar-test/logging_service/internal/configs"
	"github.com/Levap123/playstar-test/logging_service/internal/mq"
	"github.com/Levap123/playstar-test/logging_service/logs"
)

func main() {
	logger := logs.New()

	// получаем конфиги с config.yml 
	cfg, err := configs.GetConfigs()
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in getting configs")
	}

	// коннектимся к mq 
	connRabbitMQ, err := mq.InitConnection(cfg.RabbitMQ.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in connecting to rabbitMQ")
	}


}
