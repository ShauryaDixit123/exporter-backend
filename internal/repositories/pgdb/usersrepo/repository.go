//go:build wireinject
// +build wireinject

package usersrepo

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

func (r *Repository) Insert(d rdbms.CreateUserI) (string, error) {
	var id rdbms.Id
	if er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			NAME:      d.Name,
			EMAIL:     d.Email,
			PASSWORD:  d.Password,
			ROLE_ID:   d.RoleId,
			IS_PARENT: d.IsParent,
		},
	).Returning("id").Executor().ScanStructs(&id); er != nil {
		return "", er
	}
	return id.Id, nil
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
