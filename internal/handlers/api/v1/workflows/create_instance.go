package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateInstance(
	ctx *gin.Context,
) {
	var req rdbms.CreateWorkflowInstanceI
	if er := ctx.ShouldBindQuery(&req); er != nil {
		ctx.JSON(500, er)
	}
	if er := h.workflowService.CreateWorkflowInstance(req); er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, gin.H{"message": "created successfully"})
}
