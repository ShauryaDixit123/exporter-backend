package countries

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type RoutesHandler interface {
	List(ctx *gin.Context)
	Read(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Activate(ctx *gin.Context)
	Deactivate(ctx *gin.Context)
}

type Handler struct {
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
