package usersrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

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

func (r *Repository) GetById(id rdbms.Id) (rdbms.UserI, error) {
	var user rdbms.UserI
	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		NAME,
		EMAIL,
		ROLE_ID,
		PASSWORD,
		IS_PARENT,
		IS_ACTIVE,
		PRIMARY_LOCATION_ID,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.Ex{
		ID: id.Id,
	}).ScanStruct(&user); er != nil {
		return rdbms.UserI{}, er
	}
	return user, nil
}

// func (r *Repository) GetAccountOfUser
