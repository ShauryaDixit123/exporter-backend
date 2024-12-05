package v1

import (
	"exporterbackend/internal/common"
	"exporterbackend/internal/common/helper"
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"net/http"
	"strings"

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
		reqPath := ctx.Request.URL.Path
		mthd := ctx.Request.Method
		var action string
		if mthd == "POST" {
			action = strings.Join((strings.Split(reqPath, "/"))[2:], "create")
		}
		if mthd == "GET" {
			action = strings.Join((strings.Split(reqPath, "/"))[2:], "read")
		}
		if mthd == "PUT" || mthd == "PATCH" {
			action = strings.Join((strings.Split(reqPath, "/"))[2:], "update")
		}
		if mthd == "DELETE" {
			action = strings.Join((strings.Split(reqPath, "/"))[2:], "delete")
		}
		user, er := r.usersRepo.GetById(rdbms.Id{Id: id})
		if er != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": er})
		}
		if is := r.helperfunctions.CheckForPermissions(common.PermsCheck{
			RoleId: user.RoleId,
			Action: action,
		}); is {
			ctx.Next()
		}
	}
}
