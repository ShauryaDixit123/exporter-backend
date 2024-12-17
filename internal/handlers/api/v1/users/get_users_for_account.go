package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsersForAccount(
	ctx *gin.Context,
) {
	var f rdbms.GetUserForAccountReq
	if er := ctx.ShouldBindJSON(&f); er != nil {
		ctx.JSON(500, er)
		return
	}
	role, er := h.rolesRepo.GetByRole(f.Role)
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	users, er := h.userService.GetUsersForAccount(rdbms.GetUserForAccount{
		AccountId: f.AccountId,
		RoleId:    role.Id,
	})
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, users)
}
