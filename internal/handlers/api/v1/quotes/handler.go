package quotes

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger        logging.Logger
	quotesService ports.QuotesService
	imagesService ports.ImagesService
}

func NewHandler(logger logging.Logger,
	quotesService ports.QuotesService,
	imageService ports.ImagesService,
) *Handler {
	return &Handler{
		logger:        logger,
		quotesService: quotesService,
		imagesService: imageService,
	}
}

type RoutesHandler interface {
	CreateRFQ(ctx *gin.Context)
	UploadRFQItemImg(ctx *gin.Context)
}
