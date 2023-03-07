package main

import (
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
	"github.com/Levap123/playstar-test/city_service/proto"
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
