package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/google/uuid"
)

type RdbmsUsersRepository interface {
	Insert(d rdbms.CreateUserI) (uuid.UUID, error)
	GetById(id uuid.UUID) (rdbms.UserI, error)
	GetUsersForAccount(
		f rdbms.GetUserForAccount,
	) ([]rdbms.UserI, error)
}

type UsersService interface {
	Create(
		u rdbms.CreateUserRequestI,
	) (string, error)
	GetUserById(f rdbms.Id) (rdbms.GetUserResponse, error)
	GetUsersForAccount(
		f rdbms.GetUserForAccount,
	) ([]rdbms.UserI, error)
}
