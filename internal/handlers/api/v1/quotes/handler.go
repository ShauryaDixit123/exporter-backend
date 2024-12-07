package quotes

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger        logging.Logger
	quotesService ports.QuotesService
}

func NewHandler(logger logging.Logger, quotesService ports.QuotesService) *Handler {
	return &Handler{
		logger:        logger,
		quotesService: quotesService,
	}
}

type RoutesHandler interface {
	CreateRFQ(ctx *gin.Context)
}
