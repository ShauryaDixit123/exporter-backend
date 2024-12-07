package users

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger      logging.Logger
	userService ports.UsersService
}

func NewHandler(logger logging.Logger, userService ports.UsersService) *Handler {
	return &Handler{
		logger:      logger,
		userService: userService,
	}
}

type RoutesHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
}
