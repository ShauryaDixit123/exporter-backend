package workflows

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetFlowForAccount(
	ctx *gin.Context,
) {
	var r rdbms.GetFlowsForAccountI
	if er := ctx.ShouldBindJSON(&r); er != nil {
		ctx.JSON(500, er)
	}
	flow, er := h.workflowService.GetFlowForAccount(r)
	if er != nil {
		ctx.JSON(500, er)
	}
	ctx.JSON(200, flow)
}
