package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsLocationsRepository interface {
	GetUserLocations(
		f rdbms.GetUserLocationsI,
	) ([]rdbms.LocationI, error)
}
