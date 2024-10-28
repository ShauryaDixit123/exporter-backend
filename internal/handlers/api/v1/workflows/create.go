package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(ctx *gin.Context) {
	var m rdbms.CreateWorkflowI
	if er := ctx.ShouldBindJSON(&m); er != nil {
		ctx.JSON(500, er)
		return
	}
	id, er := h.workflowService.Create(m)
	if er != nil {
		ctx.JSON(500, er)
		return
	}
	ctx.JSON(200, gin.H{"id": id})
}
