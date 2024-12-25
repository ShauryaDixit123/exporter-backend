package users

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"exporterbackend/pkg/socket"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger       logging.Logger
	userService  ports.UsersService
	rolesRepo    ports.RdbmsRolesRepository
	sockekClient *map[string]*socket.Client
}

func NewHandler(logger logging.Logger,
	userService ports.UsersService,
	rolesRepo ports.RdbmsRolesRepository,
	sockekClient *map[string]*socket.Client,
) *Handler {
	return &Handler{
		logger:       logger,
		userService:  userService,
		rolesRepo:    rolesRepo,
		sockekClient: sockekClient,
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
