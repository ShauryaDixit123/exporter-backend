package accountsrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

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

func (r *Repository) Insert(d rdbms.CreateAccountI) (int, error) {
	var id rdbms.IdInt
	if er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			PRIMARY_USER_ID: d.PrimaryUserID,
			IS_ACTIVE:       d.IsActive,
		},
	).Returning("id").Executor().ScanStructs(&id); er != nil {
		return 0, er
	}
	return id.Id, nil
}

func (r *Repository) InsertAccountUser(d rdbms.CreateAccountUserI) error {
	if _, er := r.dbClient.Insert(TABLE_JOINED).Rows(
		goqu.Record{
			ACCOUNT_ID: d.AccountId,
			USER_ID:    d.UserId,
		},
	).Returning("id").Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) GetById(id int) (rdbms.AccountI, error) {
	var acc rdbms.AccountI
	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		PRIMARY_USER_ID,
		IS_ACTIVE,
		DEFAULT_WORKFLOW,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.Ex{
		ID: id,
	}).ScanStruct(&acc); er != nil {
		return rdbms.AccountI{}, er
	}
	return acc, nil
}
func (r *Repository) GetUserAccountById(id string) (rdbms.AccountI, error) {
	var acc rdbms.AccountI
	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		PRIMARY_USER_ID,
		IS_ACTIVE,
		CREATED_AT,
		MODIFIED_AT,
	).Join(goqu.I(TABLE_JOINED), goqu.On(goqu.I(PRIMARY_USER_ID).Eq(id))).ScanStruct(&acc); er != nil {
		return rdbms.AccountI{}, er
	}
	return acc, nil
}

func (r *Repository) InsertFlowInstanceAccount(
	d rdbms.CreateFlowInstanceAccountI,
) error {
	if _, er := r.dbClient.Insert(TABLE_ACCOUNT_INSTANCE).Rows(
		goqu.Record{
			ACCOUNT_ID:  d.AccountId,
			INSTANCE_ID: d.FlowInstanceId,
		},
	).Returning("id").Executor().Exec(); er != nil {
		return er
	}
	return nil
}
