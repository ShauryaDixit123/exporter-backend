package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var user rdbms.CreateUserRequestI
	if er := ctx.ShouldBindJSON(&user); er != nil {
		ctx.JSON(402, er)
	}
	id, er := h.userService.Create(user)
	if er != nil {
		ctx.JSON(402, er)
	}
	ctx.JSON(200, id)
}
