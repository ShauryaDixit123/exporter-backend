package users

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Get(ctx *gin.Context) {
	idStr := ctx.Param("id")
	if idStr == "" {
		ctx.JSON(500, "id is required")
	}
	clientMap := *h.sockekClient
	user, er := h.userService.GetUserById(rdbms.Id{Id: idStr})
	if er != nil {
		ctx.JSON(500, er)
	}
	fmt.Println(clientMap, "clientMap")
	if val, ok := clientMap[fmt.Sprintf("notification_%s", idStr)]; ok {
		val.Conn.WriteJSON(user)
	}
	ctx.JSON(200, user)
}
