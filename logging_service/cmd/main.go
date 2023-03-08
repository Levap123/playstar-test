package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Levap123/playstar-test/logging_service/internal/configs"
	"github.com/Levap123/playstar-test/logging_service/internal/mq"
	"github.com/Levap123/playstar-test/logging_service/logs"
	"github.com/rs/zerolog/log"
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

	// открываем канал с rabbitMQ
	mqChan, err := connRabbitMQ.Channel()
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in opening channel from rabbitMQ connection")
	}

	// создаем конюсмера, который будет принимать сообщения
	consumer, err := mq.NewConsumer(mqChan)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in creating consumer")
	}

	// запускаем получение сообщение с очереди в горутине
	go consumer.Consume()

	log.Info().Msg("server is ready to accept message from queue!")

	// gracefull shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	if err := mqChan.Close(); err != nil {
		logger.Fatal().Err(err).Msg("fatal in closing mq channel connection")
	}

	if err := connRabbitMQ.Close(); err != nil {
		logger.Fatal().Err(err).Msg("fatal in closing rabbit connection")

	}
}
