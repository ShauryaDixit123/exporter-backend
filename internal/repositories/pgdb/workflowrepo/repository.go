package workflowrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"
	"fmt"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/google/uuid"
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
) (uuid.UUID, error) {
	var id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE_WORKFLOW).
		Rows(goqu.Record{
			NAME:       m.Name,
			TYPE:       m.Type,
			ACCOUNT_ID: m.AccountId,
		}).Returning(ID).Executor().ScanVal(&id); er != nil {
		return uuid.Nil, er
	}
	fmt.Println("mmmfmfm")
	return id, nil
}
func (r *Repository) InsertFlow(
	f rdbms.FlowI,
) (uuid.UUID, error) {
	var id uuid.UUID

	if _, er := r.dbClient.Insert(TABLE_FLOW).Rows(goqu.Record{
		WORKFLOW_ID: f.WorkflowID,
		DESCRIPTION: f.Description,
		TYPE:        f.Type,
		ORDER:       f.Order,
		TAT:         f.TAT,
	}).Returning(ID).Executor().ScanVal(&id); er != nil {
		return uuid.Nil, er
	}
	return id, nil
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
	var id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE_FLOW_INSTANCE).Rows(f).Returning(ID).Executor().ScanVal(&id); er != nil {
		fmt.Println(er, "rmrmrm")
		return "", er
	}
	return id.String(), nil
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
) (rdbms.WorkflowI, error) {
	var wf rdbms.WorkflowI
	if _, er := r.dbClient.From(TABLE_WORKFLOW).Select(
		ID,
		NAME,
		TYPE,
		ACCOUNT_ID,
	).Where(goqu.C(ID).Eq(id)).ScanStruct(&wf); er != nil {
		return rdbms.WorkflowI{}, er
	}
	return rdbms.WorkflowI{}, nil
}

func (r *Repository) GetDetails(
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
	wid uuid.UUID,
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
	).Where(goqu.C(WORKFLOW_ID).Eq(wid.String())).Executor().ScanStructs(&ar); er != nil {
		return nil, er
	}
	return ar, nil
}

func (r *Repository) GetFlowParams(
	fpid string,
) ([]rdbms.GetFlowParamsResponseI, error) {
	var ar []rdbms.GetFlowParamsResponseI
	if er := r.dbClient.From(TABLE_FLOW_PARAMS).Select(
		ID,
		NAME,
		FLOW_ID_PARAM,
		TYPE,
		CREATED_AT,
		CREATED_AT,
		UPDATED_AT,
	).Where(goqu.C(FLOW_ID_PARAM).Eq(fpid)).ScanStructs(&ar); er != nil {
		fmt.Println("errr", er)
		return nil, er
	}
	return ar, nil
}

func (r *Repository) CreateFlowInstanceAccount(
	f rdbms.CreateFlowInstanceAccountI,
) error {
	if _, er := r.dbClient.Insert(TABLE_FLOW_INSTANCES_ACCOUNTS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) GetWorkflowByType(
	f rdbms.GetWorkflowByType,
) (rdbms.WorkflowI, error) {
	var wf rdbms.WorkflowI
	q := exp.Ex{}
	q[TYPE] = f.Type
	if f.AccountId != nil {
		q[ACCOUNT_ID] = *f.AccountId
	}
	if _, er := r.dbClient.From(TABLE_WORKFLOW).Select(
		ID,
		NAME,
		TYPE,
		ACCOUNT_ID,
	).Where(q).Executor().ScanVal(&wf); er != nil {
		return rdbms.WorkflowI{}, er
	}
	fmt.Println(wf, "dmdmdm")
	return wf, nil
}
func (r *Repository) GetFlowInstance(
	f rdbms.GetFlowInstance,
) (*rdbms.GetFlowInstanceResponseI, error) {
	var wf rdbms.GetFlowInstanceResponseI
	if _, er := r.dbClient.From(TABLE_FLOW_INSTANCE).Select(
		ID,
		WORKFLOW_ID,
		DESCRIPTION,
		TYPE,
		TITLE,
		ORDER,
		ACTIVE,
		TAT,
		INSTANCE_ID,
		STATUS,
		IS_COMPLETED,
		ASSIGNED_TO,
	).Where(
		goqu.And(
			goqu.C(INSTANCE_ID).Eq(f.InstanceId),
			goqu.C(TYPE).Eq(f.Type),
		),
	).ScanStruct(&wf); er != nil {
		return nil, er
	}
	return &wf, nil
}

func (r *Repository) GetFlowInstanceParams(
	f rdbms.GetFlowInstance,
) ([]rdbms.GetFlowInstanceParamsResponseI, error) {
	var wf []rdbms.GetFlowInstanceParamsResponseI
	ex := exp.Ex{}
	if f.Type != nil {
		ex[TYPE] = *f.Type
	}
	if f.InstanceId != nil {
		ex[INSTANCE_ID] = *f.InstanceId
	}
	if f.FlowInstanceId != nil {
		ex[FLOW_INSTANCES_ID] = *f.FlowInstanceId
	}
	if er := r.dbClient.From(TABLE_FLOW_INSTANCE).Select(
		ID,
		FLOW_INSTANCES_ID,
		NAME,
		TYPE,
		INSTANCE_PARAM_MANDATORY,
		INSTANCE_PARAM_APPROVED,
		INSTANCE_PARAM_VALUE,
	).Where(
		ex,
	).Join(goqu.T(TABLE_FLOW_INSTANCE_PARAMS), goqu.On(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ID)).Eq(
			goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, FLOW_INSTANCES_ID)),
		),
	)).ScanStructs(&wf); er != nil {
		return nil, er
	}
	return wf, nil
}
func (r *Repository) GetInstanceAccount(
	f rdbms.GetInstanceAccount,
) ([]rdbms.CreateFlowInstanceAccountI, error) {
	var wf []rdbms.CreateFlowInstanceAccountI
	if er := r.dbClient.From(TABLE_FLOW_INSTANCES_ACCOUNTS).Select(
		FLOW_INSTANCES_ID,
		ACCOUNT_ID,
		INSTANCE_ID,
		TITLE,
		DESCRIPTION,
		TYPE,
		ORDER,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ID)).As("flow_instance_id"),
		ORDER,
		TAT,
		STATUS,
		IS_COMPLETED,
		ASSIGNED_TO,
		WORKFLOW_ID,
	).Where(
		goqu.And(
			goqu.C(ACCOUNT_ID).Eq(f.AccountId),
		),
	).Join(goqu.T(TABLE_FLOW_INSTANCE), goqu.On(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCES_ACCOUNTS, INSTANCE_ID)).Eq(
			goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, INSTANCE_ID)),
		),
	)).ScanStructs(&wf); er != nil {
		return nil, er
	}
	return wf, nil
}

func (r *Repository) UpdateFlowInstanceParam(
	f rdbms.UpdateFlowInstanceParamsI,
) error {
	if _, er := r.dbClient.Update(TABLE_FLOW_INSTANCE_PARAMS).Set(
		goqu.Record{
			INSTANCE_PARAM_VALUE: f.Value,
		},
	).Where(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, ID)).Eq(f.Id),
	).Executor().Exec(); er != nil {
		return er
	}

	return nil
}

func (r *Repository) UpdateFlowInstance(
	f rdbms.UpdateFlowInstanceI,
) error {
	rec := goqu.Record{}
	if f.Status != nil {
		rec[STATUS] = *f.Status
	}
	if f.AssignedTo != nil {
		rec[ASSIGNED_TO] = *f.AssignedTo
	}
	if f.Active != nil {
		rec[ACTIVE] = *f.Active
	}
	if f.ExpiresAt != nil {
		rec[EXPIRES_AT] = *f.ExpiresAt
	}
	if _, er := r.dbClient.Update(TABLE_FLOW_INSTANCE).Set(rec).Where(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ID)).Eq(f.Id),
	).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) GetInstances(
	f rdbms.GetInstancesI,
) ([]rdbms.FlowInstanceDetails, error) {
	var wf []rdbms.FlowInstanceDetails
	q := r.dbClient.From(TABLE_FLOW_INSTANCE).Select(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ID)).As("flow_instance_id"),
		DESCRIPTION,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, TYPE)).As("flow_instance_type"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, TYPE)).As("flow_instance_params_type"),
		TITLE,
		ORDER,
		ACTIVE,
		TAT,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, INSTANCE_ID)).As("flow_instance_id"),
		STATUS,
		IS_COMPLETED,
		ASSIGNED_TO,
		WORKFLOW_ID,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, ID)).As("flow_instance_params_id"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, INSTANCE_PARAM_NAME)).As("flow_instance_params_name"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, INSTANCE_PARAM_VALUE)).As("flow_instance_params_value"),
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, INSTANCE_PARAM_MANDATORY)).As("flow_instance_params_mandatory"),
	).Join(
		goqu.I(TABLE_FLOW_INSTANCE_PARAMS), goqu.On(
			goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE_PARAMS, FLOW_INSTANCE_ID_PARAM)).Eq(
				goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ID)),
			),
		),
	).Join(goqu.I(TABLE_FLOW_INSTANCES_ACCOUNTS), goqu.On(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCES_ACCOUNTS, INSTANCE_ID)).Eq(
			goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, INSTANCE_ID)),
		),
	)).Where(
		goqu.Ex{
			fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCES_ACCOUNTS, ACCOUNT_ID): f.AccountId,
			fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, TITLE):                f.Title,
			fmt.Sprintf("%s.%s", TABLE_FLOW_INSTANCE, ORDER):                f.Order,
		},
	)
	qu, _, _ := q.ToSQL()
	fmt.Println(qu, "q")
	if er := q.Executor().ScanStructs(&wf); er != nil {
		return nil, er
	}
	return wf, nil
}
func (r *Repository) GetFlowsForAccount(
	f rdbms.GetFlowsForAccountI,
) ([]rdbms.FlowAccountsResponseI, error) {
	var wf []rdbms.FlowAccountsResponseI
	var qs string
	if f.PreOrder {
		qs = fmt.Sprintf("%s.%s", "accounts", "default_workflow_pre_order")
	} else {
		qs = fmt.Sprintf("%s.%s", "accounts", "default_workflow_post_order")
	}
	q := r.dbClient.From(TABLE_WORKFLOW).Select(
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, ID)).As("flow_id"),
		WORKFLOW_ID,
		DESCRIPTION,
		goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, TYPE)).As("flow_type"),
		TITLE,
		ORDER,
		ACTIVE,
		TAT,
	).Join(
		goqu.I(TABLE_FLOW), goqu.On(
			goqu.I(fmt.Sprintf("%s.%s", TABLE_FLOW, WORKFLOW_ID)).Eq(
				goqu.I(fmt.Sprintf("%s.%s", TABLE_WORKFLOW, ID)),
			),
		),
	).Join(
		goqu.I("accounts"),
		goqu.On(
			goqu.I(qs).Eq(
				goqu.I(fmt.Sprintf("%s.%s", TABLE_WORKFLOW, ID)),
			)),
	).Where(goqu.I(fmt.Sprintf("%s.%s", "accounts", "id")).Eq(f.AccountId))
	rqq, _, _ := q.ToSQL()
	fmt.Println(rqq, "qq")
	if er := q.
		Executor().ScanStructs(&wf); er != nil {
		return nil, er
	}
	return wf, nil
}

// "CHARGEBEE_API_KEY": "test_2Dcdw5xNBnuXOIwZhjOHPqFNrOb6O3Rax",
// "CHARGEBEE_SITE": "elixirdocs-test",
// "EMAIL_OTP": "no-reply@elixirdocs.io",
// "EMAIL_OTP_PASSWORD": "Raw58520",
// "POSTMARK_SERVER_TOKEN": "8177389f-a4b8-4a82-9d53-6d0898e7fa98",
// "POSTMARK_MESSAGE_STREAM": "unprod",
// "AWS_ACCESS_KEY_ID": "3RpXCwotXmFGCsys",
// "AWS_SECRET_ACCESS_KEY": "DOJpwMluFO3xOYjoFbpVGisQIxpliIUe",
// "AWS_BUCKET_NAME": "local-sage-documents",
// "S3_ENDPOINT": "http://127.0.0.1:9005",
// "REDIS_DB": "0",
// "REDIS_DSN": "localhost:55001",
// "REDIS_PASSWORD": "redispw",
// "ACCESS_SECRET": "HEY_BABY!",
// "REFRESH_SECRET": "YEAH_BABY!",
// "PDF_HOST": "localhost:5051",
// "NANONETS_API_KEY": "N6sJkBEGDiMvyn5NPJ2QsKv1GL3bcJJi",
// "PACKING_LIST_MODEL_ID_1": "9beed2bf-a374-42c4-9c57-78598b3d803a",
// "SHIPPING_BILL_MODEL_ID_1": "a282bed9-3990-4711-9b95-2a4af7dde904",
// "BASE_URL": "http://local.sage.com"
