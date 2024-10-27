package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
)

type RdbmsAccountsRepository interface {
	Insert(d rdbms.CreateAccountI) (int, error)
	InsertAccountUser(d rdbms.CreateAccountUserI) error
	GetById(id rdbms.Id) (rdbms.AccountI, error)
	GetUserAccountById(id rdbms.Id) (rdbms.AccountI, error)
	InsertFlowInstanceAccount(
		d rdbms.CreateFlowInstanceAccountI,
	) error
}
