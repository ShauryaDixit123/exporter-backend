package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateFlowInstanceParams(ctx *gin.Context) {
	var r rdbms.UpdateFlowInstanceParamsI
	if er := ctx.ShouldBindJSON(&r); er != nil {
		ctx.JSON(500, er)
	}
	if er := h.workflowService.UpdateFlowInstanceParam(r); er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, gin.H{"message": "updated successfully"})
}
