package helper

import (
	"exporterbackend/internal/common"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
)

type HelperFunctions interface {
	CheckForPermissions(p common.PermissionCheck) bool
	ParseURLAndAction(url, method string) string
}

type HelperRepository struct {
	logger    logging.Logger
	rolesRepo ports.RdbmsRolesRepository
}

func NewHelperRepository(
	logger logging.Logger,
	rolesRepo ports.RdbmsRolesRepository,
) *HelperRepository {
	return &HelperRepository{
		logger:    logger,
		rolesRepo: rolesRepo,
	}
}
