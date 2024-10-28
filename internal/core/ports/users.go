package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"

	"github.com/google/uuid"
)

type RdbmsUsersRepository interface {
	Insert(d rdbms.CreateUserI) (uuid.UUID, error)
	GetById(id rdbms.Id) (rdbms.UserI, error)
}

type UsersService interface {
	Create(
		u rdbms.CreateUserI,
	) (string, error)
}
