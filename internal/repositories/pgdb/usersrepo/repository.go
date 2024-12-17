package usersrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"
	"fmt"

	"github.com/doug-martin/goqu/v9"
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

func (r *Repository) Insert(d rdbms.CreateUserI) (uuid.UUID, error) {
	var id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			NAME:                d.Name,
			EMAIL:               d.Email,
			PASSWORD:            d.Password,
			ROLE_ID:             d.RoleId,
			IS_PARENT:           d.IsParent,
			PRIMARY_LOCATION_ID: d.PrimaryLocationID,
		},
	).Returning(ID).Executor().ScanVal(&id); er != nil {
		return uuid.UUID{}, er
	}
	return id, nil
}

func (r *Repository) GetById(id uuid.UUID) (rdbms.UserI, error) {
	var user rdbms.UserI
	fmt.Println(id, "ududud")
	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		NAME,
		EMAIL,
		ROLE_ID,
		PASSWORD,
		IS_PARENT,
		ACCESS_TOKEN,
		IS_ACTIVE,
		PRIMARY_LOCATION_ID,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.Ex{
		ID: id,
	}).ScanStruct(&user); er != nil {
		return rdbms.UserI{}, er
	}
	return user, nil
}

func (r *Repository) GetUsersForAccount(
	f rdbms.GetUserForAccount,
) ([]rdbms.UserI, error) {
	var rs []rdbms.UserI
	q := r.dbClient.From("accounts_users_map").Select(
		goqu.I(fmt.Sprintf("%s.%s", TABLE, ID)).As(ID),
		NAME,
		EMAIL,
		ROLE_ID,
		IS_PARENT,
		IS_ACTIVE,
		PRIMARY_LOCATION_ID,
	).Join(
		goqu.T(TABLE), goqu.On(
			goqu.I(fmt.Sprintf("%s.%s", "accounts_users_map", USER_ID)).Eq(goqu.I(fmt.Sprintf("%s.%s", TABLE, ID))),
		),
	).Where(
		goqu.Ex{
			fmt.Sprintf("%s.%s", "accounts_users_map", ACCOUNT_ID): f.AccountId,
			fmt.Sprintf("%s.%s", TABLE, "role_id"):                 f.RoleId,
		},
	)
	if er := q.Executor().ScanStructs(&rs); er != nil {
		return []rdbms.UserI{}, er
	}
	return rs, nil
}

// func (r *Repository) GetAccountOfUser
