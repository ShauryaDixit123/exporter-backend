package locationsrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger   logging.Logger
	dbClient *goqu.Database
}

func New(logger logging.Logger,
	dbClient *goqu.Database) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r *Repository) Insert(d rdbms.CreateLocationI) (string, error) {
	var id rdbms.Id
	if er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			LINE1:      d.Line1,
			LINE2:      d.Line2,
			AREA:       d.Area,
			CITY:       d.City,
			STATE:      d.State,
			COUNTRY_ID: d.CountryID,
			PINCODE:    d.Pincode,
		},
	).Returning("id").Executor().ScanStructs(&id); er != nil {
		return "", er
	}
	return id.Id, nil
}

func (r *Repository) GetById(id rdbms.Id) (rdbms.LocationI, error) {
	var user rdbms.LocationI

	if _, er := r.dbClient.From(TABLE).Select(
		ID,
		LINE1,
		LINE2,
		AREA,
		CITY,
		STATE,
		COUNTRY_ID,
		PINCODE,
		CREATED_AT,
		MODIFIED_AT,
	).Where(goqu.Ex{
		ID: id,
	}).ScanStruct(&user); er != nil {
		return rdbms.LocationI{}, er
	}
	return user, nil
}
