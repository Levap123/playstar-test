package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Levap123/playstar-test/logging_service/internal/configs"
	"github.com/Levap123/playstar-test/logging_service/internal/mq"
	"github.com/Levap123/playstar-test/logging_service/internal/repository/postgres"
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

	// создаем коннекшн к бд
	DB, err := postgres.InitDB(cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in initalizating DB")

	}

	ctxPing, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if err := DB.PingContext(ctxPing); err != nil {
		logger.Fatal().Err(err).Msg("fatal in pinging DB")

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

	// репозиторий для хранения логов
	repo := postgres.NewLogsRepo(DB, logger)

	// создаем конюсмера, который будет принимать сообщения
	consumer, err := mq.NewConsumer(mqChan, logger, repo)
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

	logger.Info().Msg("server stopped")
}
