package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/google/uuid"
)

type RdbmsWorkflowRepository interface {
	Insert(
		m rdbms.WorkflowI,
	) (uuid.UUID, error)
	GetDetails(
		id string,
	) ([]rdbms.GetWorkflowI, error)
	GetAll(
		of string,
	) ([]rdbms.WorkflowI, error)
	InsertFlow(
		f rdbms.FlowI,
	) (uuid.UUID, error)
	InsertFlowParams(
		f []rdbms.FlowParamI,
	) error
	InsertFlowInstance(
		f rdbms.FlowInstanceI,
	) (string, error)
	InsertFlowInstanceParam(
		f []rdbms.FlowInstanceParamI,
	) error
	GetFlows(
		wid uuid.UUID,
	) ([]rdbms.FlowI, error)
	GetFlowParams(
		fpid string,
	) ([]rdbms.GetFlowParamsResponseI, error)
	CreateFlowInstanceAccount(
		f rdbms.CreateFlowInstanceAccountI,
	) error
	GetWorkflowByType(
		f rdbms.GetWorkflowByType,
	) (rdbms.WorkflowI, error)
	Get(
		id string,
	) (rdbms.WorkflowI, error)
	GetFlowInstance(
		f rdbms.GetFlowInstance,
	) (*rdbms.GetFlowInstanceResponseI, error)
	GetFlowInstanceParams(
		f rdbms.GetFlowInstance,
	) (*rdbms.GetFlowInstanceParamsResponseI, error)
}

type WorkflowService interface {
	Create(m rdbms.CreateWorkflowI) (string, error)
	CreateWorkflowInstance(
		m rdbms.CreateWorkflowInstanceI,
	) (string, error)
	Get(id string) ([]rdbms.GetWorkflowI, error)
	GetAll(of string) ([]rdbms.WorkflowI, error)
	AttachToWorkflow(
		d rdbms.AttachWorkflowReqI,
	) (rdbms.AttachWorkflowI, error)
}
