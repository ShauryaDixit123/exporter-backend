package workflowrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger   logging.Logger
	dbClient *goqu.Database
}

func New(logger logging.Logger,
	dbClient *goqu.Database) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}
func (r *Repository) Insert(
	m rdbms.WorkflowI,
) (string, error) {
	var id rdbms.Id
	if _, er := r.dbClient.Insert(TABLE_WORKFLOW).
		Rows(m).Returning("id").Executor().ScanStruct(&id); er != nil {
		return "", er
	}
	return id.Id, nil
}
func (r *Repository) InsertFlow(
	f rdbms.FlowI,
) (string, error) {
	var id rdbms.Id
	if _, er := r.dbClient.Insert(TABLE_FLOW).Rows(f).Returning("id").Executor().ScanStruct(&id); er != nil {
		return "", er
	}
	return id.Id, nil
}
func (r *Repository) InsertFlowParams(
	f []rdbms.FlowParamI,
) error {
	if _, er := r.dbClient.Insert(TABLE_FLOW_PARAMS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}
func (r *Repository) InsertFlowInstance(
	f rdbms.FlowInstanceI,
) (string, error) {
	var id string
	if _, er := r.dbClient.Insert(TABLE_FLOW_INSTANCE).Rows(f).Returning("id").Executor().ScanStruct(&id); er != nil {
		return "", er
	}
	return id, nil
}
func (r *Repository) InsertFlowInstanceParam(
	f []rdbms.FlowInstanceParamI,
) error {
	if _, er := r.dbClient.Insert(TABLE_FLOW_INSTANCE_PARAMS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) Get(
	id string,
) ([]rdbms.GetWorkflowI, error) {
	var wf []rdbms.GetWorkflowI
	if er := r.dbClient.From(TABLE_WORKFLOW).Select(
		ID,
		NAME,
		TYPE,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, ID)).As("flow_id"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, DESCRIPTION)),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, TYPE)),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, ORDER)),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, ACTIVE)),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, TAT)),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_PARAMS, ID)).As("flow_params_id"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_PARAMS, PARAM_NAME)).As("flow_params_name"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_PARAMS, PARAM_TYPE)).As("flow_params_type"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_PARAMS, PARAM_MANDATORY)),
	).Join(goqu.I(TABLE_FLOW), goqu.On(goqu.I(fmt.Sprintf(
		"%s.%s", TABLE_FLOW, WORKFLOW_ID,
	),
	).Eq(fmt.Sprintf(
		"%s.%s", TABLE_WORKFLOW, ID)))).Join(
		goqu.I(TABLE_FLOW_PARAMS), goqu.On(goqu.I(fmt.Sprintf(
			"%s.%s", TABLE_FLOW_PARAMS, FLOW_ID),
		).Eq(fmt.Sprintf(
			"%s.%s", TABLE_FLOW, ID))),
	).ScanStructs(wf); er != nil {
		return nil, er
	}
	return wf, nil
}

func (r *Repository) GetAll(
	of string,
) ([]rdbms.WorkflowI, error) {
	var vals []rdbms.WorkflowI
	if er := r.dbClient.From(TABLE_WORKFLOW).Select("*").ScanStructs(&vals); er != nil {
		return nil, er
	}
	return vals, nil
}

func (r *Repository) GetFlows(
	wid string,
) ([]rdbms.FlowI, error) {
	var ar []rdbms.FlowI
	if er := r.dbClient.From(TABLE_FLOW).Select(
		ID,
		WORKFLOW_ID,
		DESCRIPTION,
		TYPE,
		TITLE,
		ORDER,
		ACTIVE,
		TAT,
	).Where(goqu.C(WORKFLOW_ID).Eq(wid)).ScanStructs(&ar); er != nil {
		return nil, er
	}
	return ar, nil
}

func (r *Repository) GetFlowParams(
	fpid string,
) ([]rdbms.FlowParamI, error) {
	var ar []rdbms.FlowParamI
	if er := r.dbClient.From(TABLE_FLOW_PARAMS).Select(
		ID,
		FLOW_ID,
		NAME,
		TYPE,
		CREATED_AT,
		CREATED_AT,
		UPDATED_AT,
	).Where(goqu.C(WORKFLOW_ID).Eq(fpid)).ScanStructs(&ar); er != nil {
		return nil, er
	}
	return ar, nil
}
