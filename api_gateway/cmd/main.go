package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	apiclients "github.com/Levap123/playstar-test/api_gateway/internal/api_clients"
	"github.com/Levap123/playstar-test/api_gateway/internal/configs"
	"github.com/Levap123/playstar-test/api_gateway/internal/handler"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := configs.GetConfigs()
	if err != nil {
		log.Fatal(err)
	}

	// контекст для подключения к сервису получения города
	ctxWeather, cityCancel := context.WithTimeout(context.Background(), time.Second)
	defer cityCancel()

	// dev затычка
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	// "коннектимся" к сервису, адрес получаем с конфигов
	connWeathersrv, err := grpc.DialContext(ctxWeather, cfg.WeatherService.Addr, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer connWeathersrv.Close()

	weatherClient := apiclients.NewWeatherClient(connWeathersrv)

	handler := handler.NewHandler(weatherClient)

	log.Println("server started")
	go func() {
		if err := http.ListenAndServe(cfg.Server.Addr, handler.InitRoutes()); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	if err := connWeathersrv.Close(); err != nil {
		log.Fatal(err)
	}
}
