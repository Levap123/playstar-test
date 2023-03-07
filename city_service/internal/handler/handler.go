package handler

import (
	"context"
	"errors"
	"time"

	apiclients "github.com/Levap123/playstar-test/city_service/internal/api_clients"
	"github.com/Levap123/playstar-test/city_service/internal/logs"
	"github.com/Levap123/playstar-test/city_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CityHandler struct {
	proto.UnimplementedCityServiceServer
	cc     *apiclients.CoordinatesClient
	logger *logs.Logger
}

func NewCityHandler(cc *apiclients.CoordinatesClient, logger *logs.Logger) *CityHandler {
	return &CityHandler{
		cc:     cc,
		logger: logger,
	}
}

func (h *CityHandler) GetCity(ctx context.Context, req *proto.GetCityRequest) (*proto.GetCityResponse, error) {
	h.logger.Debug().Msg("get city handler")

	reqCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	city, err := h.cc.GetCityFromCoordiantes(reqCtx, req.Latitude, req.Longitude)
	if err != nil {
		h.logger.Err(err).Msg("something went wrong in get city handler")

		switch {
		case errors.Is(err, apiclients.ErrBadRequest):
			return nil, status.Error(codes.InvalidArgument, apiclients.ErrBadRequest.Error())
		case errors.Is(err, apiclients.ErrInternal):
			return nil, status.Error(codes.Internal, apiclients.ErrInternal.Error())
		default:
			return nil, err
		}
	}

	h.logger.Info().Msg("success in sending response")
	return &proto.GetCityResponse{
		City: city.City,
	}, nil
}
