package rolesrepo

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

func New(
	logger logging.Logger,
	dbClient *goqu.Database,
) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r *Repository) GetById(id rdbms.Id) (rdbms.RoleI, error) {
	var role rdbms.RoleI
	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		ROLE,
		DISPLAY_VALUE,
	).Where(goqu.Ex{
		ID: id,
	}).Executor().ScanStruct(&role); er != nil {
		return rdbms.RoleI{}, er
	}
	return role, nil
}

func (r *Repository) GetByRole(role string) (rdbms.RoleI, error) {
	var rp rdbms.RoleI
	q := r.dbClient.From(TABLE).Select(
		ID,
		ROLE,
		DISPLAY_VALUE,
	).Where(goqu.Ex{
		ROLE: role,
	})
	qw, _, er := q.ToSQL()
	if er != nil {
		return rdbms.RoleI{}, er
	}
	fmt.Println(qw, "eeemm")
	if _, er := q.Executor().ScanStruct(&rp); er != nil {
		return rdbms.RoleI{}, er
	}
	fmt.Println(rp)
	return rp, nil
}
