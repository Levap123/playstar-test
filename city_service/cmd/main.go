package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiclients "github.com/Levap123/playstar-test/city_service/internal/api_clients"
	"github.com/Levap123/playstar-test/city_service/internal/configs"
	"github.com/Levap123/playstar-test/city_service/internal/handler"
	"github.com/Levap123/playstar-test/city_service/internal/logs"
	"github.com/Levap123/playstar-test/city_service/internal/producer"
	"github.com/Levap123/playstar-test/city_service/proto"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger, err := logs.New()
	if err != nil {
		log.Fatal(err)
	}

	// получаем конфиги с config.yml файла в структуру
	cfg, err := configs.GetConfigs()
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in getting configs")
	}

	// подключаемся к rabbitmq серверу
	connRabbitMQ, err := amqp.Dial(cfg.RabbitMQ.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in connecting to rabbitmq")
	}

	// открываем канал
	channelRabbitMQ, err := connRabbitMQ.Channel()
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in openning channel from rabbitmq")
	}

	// создаем инстанс, который будет отправлять сообщения в mq
	producer, err := producer.NewProducer(channelRabbitMQ, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in creating producer")

	}

	// будет отправлять сообщения в mq
	go producer.Produce()

	// создаем клиента для отправки реквество на стороннее апи
	client := &http.Client{
		Timeout: time.Duration(cfg.Client.TimeoutSeconds) * time.Second,
	}

	// враппер для килента
	coordCl := apiclients.NewCoordinatesClient(client)

	// хендлер который реализует gRPC сервис
	handler := handler.NewCityHandler(coordCl, logger)

	// берет адрес с кофингов, на котором будет слушать реквесты
	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in listen")
	}
	defer listener.Close()

	srv := grpc.NewServer()
	proto.RegisterCityServiceServer(srv, handler)
	reflection.Register(srv)

	// graceful shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	logger.Info().Msgf("server started on: %s", cfg.Server.Addr)
	go func() {
		if err := srv.Serve(listener); err != nil {
			logger.Fatal().Err(err).Msg("fatal in starting server")
		}
	}()

	<-quit

	srv.GracefulStop()

	logger.Info().Msg("server stopped")
}
