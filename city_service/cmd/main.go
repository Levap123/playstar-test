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
	"github.com/Levap123/playstar-test/city_service/internal/city"
	"github.com/Levap123/playstar-test/city_service/internal/configs"
	"github.com/Levap123/playstar-test/city_service/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg, err := configs.GetConfigs()
	if err != nil {
		log.Fatalf("fatal in getting configs: %v", err)
	}

	client := &http.Client{
		Timeout: time.Duration(cfg.Client.TimeoutSeconds) * time.Second,
	}

	coordCl := apiclients.NewCoordinatesClient(client)

	handler := city.NewCityHandler(coordCl)

	listener, err := net.Listen("tcp", cfg.Server.Addr)
	if err != nil {
		log.Fatalf("error in listen: %v", err)
	}
	defer listener.Close()

	srv := grpc.NewServer()
	proto.RegisterCityServiceServer(srv, handler)
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
