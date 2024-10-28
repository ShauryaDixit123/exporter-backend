package orders

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var d rdbms.CreatePurchaseOrder
	if er := ctx.ShouldBindJSON(&d); er != nil {
		ctx.JSON(500, er)
		return
	}
	er := h.ordersService.CreatePurchaseOrder(d)
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, gin.H{"message": "created_successfully"})
}
