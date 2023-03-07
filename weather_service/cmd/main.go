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
	"github.com/Levap123/playstar-test/weather_service/internal/weather"
	"github.com/Levap123/playstar-test/weather_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := configs.GetConfigs()
	if err != nil {
		log.Fatalf("fatal in getting configs: %v", err)
	}

	ctxCity, cityCancel := context.WithTimeout(context.Background(), time.Second)
	defer cityCancel()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	connCitysrv, err := grpc.DialContext(ctxCity, cfg.CityService.Addr, opts...)
	if err != nil {
		log.Fatalf("fatal in connect to city service: %v", err)
	}
	defer connCitysrv.Close()

	wclient := &http.Client{
		Timeout: time.Duration(cfg.Client.TimeoutSeconds) * time.Second,
	}

	weatherCl := apiclients.NewWeatherClient(cfg.WeatherApi.Key, wclient)
	cityCl := apiclients.NewCityClient(connCitysrv)

	handler := weather.NewWeatherHandler(cityCl, weatherCl)

	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatalf("error in listen: %v", err)
	}
	defer listener.Close()

	srv := grpc.NewServer()
	proto.RegisterWeatherServiceServer(srv, handler)
	reflection.Register(srv)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	log.Printf("server started on: %s", cfg.Server.Addr)
	go func() {
		if err := srv.Serve(listener); err != nil {
			log.Fatalf("error in starting serving: %v", err)
		}
	}()

	<-quit

	srv.GracefulStop()

	log.Println("server stopped")
}
