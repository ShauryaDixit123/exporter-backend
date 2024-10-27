package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var m rdbms.CreateWorkflowI
	if er := ctx.ShouldBindJSON(&m); er != nil {
		ctx.JSON(500, er)
	}
	if er := h.workflowService.Create(m); er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, gin.H{"message": "created successfully"})
}
