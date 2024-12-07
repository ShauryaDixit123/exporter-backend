package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(500, "id is required")
	}
	user, er := h.userService.GetUserById(rdbms.Id{Id: idStr})
	if er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, user)
}
