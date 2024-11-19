package workflows

import (
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	logger          logging.Logger
	workflowService ports.WorkflowService
}

func NewHandler(logger logging.Logger, workflowService ports.WorkflowService) *Handler {
	return &Handler{
		logger:          logger,
		workflowService: workflowService,
	}
}

type RoutesHandler interface {
	Create(ctx *gin.Context)
	CreateInstance(
		ctx *gin.Context,
	)
	UpdateFlowInstance(ctx *gin.Context)
	UpdateFlowInstanceParams(ctx *gin.Context)
	GetInstances(ctx *gin.Context)
	GetFlowForAccount(
		ctx *gin.Context,
	)
}
