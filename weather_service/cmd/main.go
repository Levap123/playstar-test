package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiclients "github.com/Levap123/playstar-test/weather_service/internal/api_clients"
	"github.com/Levap123/playstar-test/weather_service/internal/configs"
	"github.com/Levap123/playstar-test/weather_service/internal/handler"
	"github.com/Levap123/playstar-test/weather_service/internal/logs"
	"github.com/Levap123/playstar-test/weather_service/internal/producer"
	"github.com/Levap123/playstar-test/weather_service/internal/validator"
	"github.com/Levap123/playstar-test/weather_service/proto"
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

	// контекст для подключения к сервису получения города
	ctxCity, cityCancel := context.WithTimeout(context.Background(), time.Second)
	defer cityCancel()

	// dev затычка
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// "коннектимся" к сервису, адрес получаем с конфигов
	connCitysrv, err := grpc.DialContext(ctxCity, cfg.CityService.Addr, opts...)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in connecting to the city service")
	}
	defer connCitysrv.Close()

	// подключаемся к rabbitmq серверу
	connRabbitMQ, err := amqp.Dial(cfg.RabbitMQ.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in connecting to rabbitmq")
	}

	// открываем канал
	mqChan, err := connRabbitMQ.Channel()
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in openning channel from rabbitmq")
	}

	// создаем инстанс, который будет отправлять сообщения в mq
	producer, err := producer.NewProducer(mqChan, logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in creating producer")

	}

	// будет отправлять сообщения в mq
	go producer.Produce()

	// клиент для отправления реквестов на сторонее апи
	wclient := &http.Client{
		Timeout: time.Duration(cfg.Client.TimeoutSeconds) * time.Second,
	}

	// врапперы http и gRPC клиентов
	weatherCl := apiclients.NewWeatherClient(cfg.WeatherApi.Key, wclient)
	cityCl := apiclients.NewCityClient(connCitysrv)

	// валидируем координаты, которые получаем от юзера
	validator := validator.New()

	// хендлер который реализует gRPC сервис
	handler := handler.NewWeatherHandler(cityCl, weatherCl, validator, logger)

	// берет адрес с кофингов, на котором будет слушать реквесты
	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		logger.Fatal().Err(err).Msg("fatal in listening")
	}
	defer listener.Close()

	// создаем gRPC сервер
	srv := grpc.NewServer()
	proto.RegisterWeatherServiceServer(srv, handler)
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

	if err := mqChan.Close(); err != nil {
		logger.Fatal().Err(err).Msg("fatal in closing mq channel connection")
	}

	if err := connRabbitMQ.Close(); err != nil {
		logger.Fatal().Err(err).Msg("fatal in closing rabbit connection")

	}

	logger.Info().Msg("server stopped")
}
