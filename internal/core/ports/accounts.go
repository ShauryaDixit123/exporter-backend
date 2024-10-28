package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
)

type RdbmsAccountsRepository interface {
	Insert(d rdbms.CreateAccountI) (int, error)
	InsertAccountUser(d rdbms.CreateAccountUserI) error
	GetById(id int) (rdbms.AccountI, error)
	GetUserAccountById(id string) (rdbms.AccountI, error)
	InsertFlowInstanceAccount(
		d rdbms.CreateFlowInstanceAccountI,
	) error
}
