package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsUsersRepository interface {
	Insert(d rdbms.CreateUserI) (string, error)
	GetById(id rdbms.Id) (rdbms.UserI, error)
}

type UsersService interface {
	Create(
		u rdbms.CreateUserI,
	) (string, error)
}
