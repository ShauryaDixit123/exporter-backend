package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersForAccount(
	ctx *gin.Context,
) {
	var f rdbms.GetUserForAccount
	if er := ctx.ShouldBindJSON(&f); er != nil {
		ctx.JSON(500, er)
		return
	}
	users, er := h.userService.GetUsersForAccount(f)
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, users)
}
