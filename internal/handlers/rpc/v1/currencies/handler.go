package currencies

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	pb "exporterbackend/proto/gen/app/v1"
)

type Handler struct {
	pb.UnimplementedCurrenciesServiceServer
	logger            logging.Logger
	currenciesService ports.CurrenciesService
}

func NewHandler(
	logger logging.Logger,
	currenciesService ports.CurrenciesService,
) *Handler {
	return &Handler{
		logger:            logger,
		currenciesService: currenciesService,
	}
}
