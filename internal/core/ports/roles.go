package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsRolesRepository interface {
	GetById(id rdbms.Id) (rdbms.RoleI, error)
	GetByRole(role string) (rdbms.RoleI, error)
}
