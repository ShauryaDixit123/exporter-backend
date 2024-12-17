package users

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger      logging.Logger
	userService ports.UsersService
	rolesRepo   ports.RdbmsRolesRepository
}

func NewHandler(logger logging.Logger, userService ports.UsersService, rolesRepo ports.RdbmsRolesRepository) *Handler {
	return &Handler{
		logger:      logger,
		userService: userService,
		rolesRepo:   rolesRepo,
	}
}

type RoutesHandler interface {
	Create(ctx *gin.Context)
	Get(ctx *gin.Context)
	GetUsersForAccount(
		ctx *gin.Context,
	)
	GetLocations(
		ctx *gin.Context,
	)
}
