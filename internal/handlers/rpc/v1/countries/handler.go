package countries

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	pb "exporterbackend/proto/gen/app/v1"
)

type Handler struct {
	pb.UnimplementedCountriesServiceServer
	logger           logging.Logger
	countriesService ports.CountriesService
}

func NewHandler(
	logger logging.Logger,
	countriesService ports.CountriesService,
) *Handler {
	return &Handler{
		logger:           logger,
		countriesService: countriesService,
	}
}
