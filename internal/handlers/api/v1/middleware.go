package v1

import (
	"exporterbackend/internal/common"
	"exporterbackend/internal/common/helper"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RouteMiddleware struct {
	helperfunctions helper.HelperFunctions
	usersRepo       ports.RdbmsUsersRepository
	logger          logging.Logger
}

type RouteMiddlewares interface {
	PermissionsMiddleware() gin.HandlerFunc
}

func NewMiddleware(
	logger logging.Logger,
	helperfunctions helper.HelperFunctions,
	usersRepo ports.RdbmsUsersRepository,
) *RouteMiddleware {
	return &RouteMiddleware{
		helperfunctions: helperfunctions,
		usersRepo:       usersRepo,
		logger:          logger,
	}
}

func (r *RouteMiddleware) PermissionsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Request.Header.Get("id")
		act := ctx.Request.Header.Get("action")
		if act == "" {
			act = ctx.Request.Method
		}
		action := r.helperfunctions.ParseURLAndAction(ctx.Request.URL.Path, act)
		uid, er := uuid.Parse(id)
		if er != nil {
			r.logger.Error(
				"INVALID_USER_ID",
				"error in parsing to uuid",
				er,
				map[string]any{
					"user_id":        uid,
					"user_id_string": id,
				},
				map[string]any{},
			)
		}
		user, er := r.usersRepo.GetById(uid)
		if er != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": er})
		}
		if is := r.helperfunctions.CheckForPermissions(common.PermissionCheck{
			RoleId: user.RoleId,
			Action: action,
		}); is {
			ctx.Next()
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		}
	}
}
