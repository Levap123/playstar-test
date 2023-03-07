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
	"github.com/Levap123/playstar-test/weather_service/internal/validator"
	"github.com/Levap123/playstar-test/weather_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	logger := logs.New()

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

	log.Printf("server started on: %s", cfg.Server.Addr)
	go func() {
		if err := srv.Serve(listener); err != nil {
			logger.Fatal().Err(err).Msg("fatal in starting server")
		}
	}()

	<-quit

	srv.GracefulStop()

	logger.Info().Msg("server stopped")
}
