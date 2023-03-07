package city

import "github.com/Levap123/playstar-test/city_service/proto"

type Coordinates struct {
	Latitude  float32
	Longitude float32
}

type City struct {
	City string `json:"city,omitempty"`
}

func FromProtoRequestToCoordinates(req *proto.GetCityRequest) Coordinates {
	return Coordinates{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
