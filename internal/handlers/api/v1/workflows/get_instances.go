package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetInstances(ctx *gin.Context) {
	var r rdbms.GetInstancesI
	if er := ctx.ShouldBindJSON(&r); er != nil {
		ctx.JSON(500, er)
	}
	instances, er := h.workflowService.GetInstances(r)
	if er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, instances)
}
