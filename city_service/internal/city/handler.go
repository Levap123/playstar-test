package city

import (
	"context"
	"errors"
	"time"

	apiclients "github.com/Levap123/playstar-test/city_service/internal/api_clients"
	"github.com/Levap123/playstar-test/city_service/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CityHadnler struct {
	proto.UnimplementedCityServiceServer
	cc *apiclients.CoordinatesClient
}

func (h *CityHadnler) GetCity(ctx context.Context, req *proto.GetCityRequest) (*proto.GetCityResponse, error) {
	reqCtx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	city, err := h.cc.GetCityFromCoordiantes(reqCtx, req.Latitude, req.Longitude)
	if err != nil {
		switch {
		case errors.Is(err, apiclients.ErrBadRequest):
			return nil, status.Error(codes.InvalidArgument, apiclients.ErrBadRequest.Error())
		case errors.Is(err, apiclients.ErrInternal):
			return nil, status.Error(codes.Internal, apiclients.ErrInternal.Error())
		default:
			return nil, err
		}
	}
	return &proto.GetCityResponse{
		City: city.City,
	}, nil
}
