package orders

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger        logging.Logger
	ordersService ports.OrdersService
}

func NewHandler(logger logging.Logger, ordersService ports.OrdersService) *Handler {
	return &Handler{
		logger:        logger,
		ordersService: ordersService,
	}
}

type RoutesHandler interface {
	Create(ctx *gin.Context)
}
