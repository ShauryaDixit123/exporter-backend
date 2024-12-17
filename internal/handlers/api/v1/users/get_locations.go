package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetLocations(ctx *gin.Context) {
	var f rdbms.GetUserLocationsI
	if er := ctx.ShouldBindJSON(&f); er != nil {
		ctx.JSON(500, er)
		return
	}
	locations, er := h.userService.GetLocationsForUser(f)
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, locations)
}
