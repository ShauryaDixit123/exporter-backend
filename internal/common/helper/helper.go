package helper

import (
	"exporterbackend/internal/common"
	"exporterbackend/pkg/logging"
)

type HelperFunctions interface {
	CheckForPermissions(p common.PermissionCheck) bool
	ParseURLAndAction(url, method string) string
}

type HelperRepository struct {
	logger logging.Logger
}

func NewHelperRepository(
	logger logging.Logger,
) *HelperRepository {
	return &HelperRepository{
		logger: logger,
	}
}
