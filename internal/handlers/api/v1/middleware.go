package v1

import (
	"exporterbackend/internal/common"
	"exporterbackend/internal/common/helper"
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RouteMiddleware struct {
	helperfunctions helper.HelperFunctions
	usersRepo       ports.RdbmsUsersRepository
}

type RouteMiddlewares interface {
	PermissionsMiddleware() gin.HandlerFunc
}

func NewMiddleware(
	helperfunctions helper.HelperFunctions,
	usersRepo ports.RdbmsUsersRepository,
) *RouteMiddleware {
	return &RouteMiddleware{
		helperfunctions: helperfunctions,
		usersRepo:       usersRepo,
	}
}

func (r *RouteMiddleware) PermissionsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Request.Header.Get("id")
		action := r.helperfunctions.ParseURLAndAction(ctx.Request.URL.Path, ctx.Request.Method)
		user, er := r.usersRepo.GetById(rdbms.Id{Id: id})
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
