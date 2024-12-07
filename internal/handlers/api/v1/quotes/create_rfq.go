package quotes

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRFQ(ctx *gin.Context) {
	var m rdbms.CreateRFQRequestI
	if er := ctx.ShouldBindJSON(&m); er != nil {
		ctx.JSON(500, er)
		return
	}
	if er := h.quotesService.CreateRFQ(m); er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, gin.H{"message": "created successfully"})
}
