package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsWorkflowRepository interface {
	Insert(
		m rdbms.WorkflowI,
	) (string, error)
	Get(
		id string,
	) ([]rdbms.GetWorkflowI, error)
	GetAll(
		of string,
	) ([]rdbms.WorkflowI, error)
	InsertFlow(
		f rdbms.FlowI,
	) (string, error)
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
		wid string,
	) ([]rdbms.FlowI, error)
	GetFlowParams(
		fpid string,
	) ([]rdbms.FlowParamI, error)
}

type WorkflowService interface {
	Create(m rdbms.CreateWorkflowI) error
	CreateWorkflowInstance(
		m rdbms.CreateWorkflowInstanceI,
	) error
	Get(id string) ([]rdbms.GetWorkflowI, error)
	GetAll(of string) ([]rdbms.WorkflowI, error)
}
