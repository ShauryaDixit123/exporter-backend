package countriesrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
)

type Repository struct {
	logger logging.Logger
	goquDB *goqu.Database
}

func New(
	logger logging.Logger,
	goquDB *goqu.Database,
) *Repository {
	return &Repository{
		logger: logger,
		goquDB: goquDB,
	}
}

func (r *Repository) SelectOne(
	isoCode string,
) (rdbms.Country, bool, error) {
	var country rdbms.Country

	exists, err := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ISO_CODE,
		NAME,
		ISO_3,
		CALLING_CODE,
		IS_ACTIVE,
		CREATED_AT,
		MODIFIED_AT,
	).Where(
		goqu.C(ISO_CODE).Eq(isoCode),
	).ScanStruct(&country)

	if err != nil {
		return rdbms.Country{}, false, err
	}

	return country, exists, nil
}

func (r *Repository) SelectMany(
	limit uint,
	offset uint,
	search ports.CountriesSearch,
	sort ports.CountriesSort,
	filters ports.CountriesFilters,
) ([]rdbms.Country, error) {

	var countries []rdbms.Country
	w := []exp.Expression{}

	if filters.IsActive != nil {
		w = append(w, goqu.C(IS_ACTIVE).Eq(*filters.IsActive))
	}

	searchClauses := r.buildSearchWhereClauses(search)

	if len(searchClauses) > 1 {
		w = append(w, goqu.Or(searchClauses...))
	} else if len(searchClauses) == 1 {
		w = append(w, searchClauses[0])
	}

	query := r.goquDB.From(
		TABLE,
	).Prepared(true).Select(
		ISO_CODE,
		NAME,
		ISO_3,
		CALLING_CODE,
		IS_ACTIVE,
		CREATED_AT,
		MODIFIED_AT,
	).Where(w...).Order(
		r.getOrderedExpression(sort),
	).Limit(
		limit,
	).Offset(
		offset,
	)

	err := query.ScanStructs(&countries)

	if err != nil {
		return []rdbms.Country{}, err
	}

	return countries, nil
}

func (r *Repository) InsertOne(
	isoCode string,
	name string,
	iso3 *string,
	callingCode uint,
	isActive bool,
) error {
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			ISO_CODE:     isoCode,
			NAME:         name,
			ISO_3:        iso3,
			CALLING_CODE: callingCode,
			IS_ACTIVE:    isActive,
		},
	).Executor().Exec()

	return err
}

func (r *Repository) UpdateOne(
	isoCode string,
	name *string,
	iso3 *string,
	callingCode *uint,
	isActive *bool,
) (int64, error) {
	updates := goqu.Record{}

	if name != nil {
		updates[NAME] = *name
	}

	if iso3 != nil {
		updates[ISO_3] = *iso3
	}

	if callingCode != nil {
		updates[CALLING_CODE] = *callingCode
	}

	if isActive != nil {
		updates[IS_ACTIVE] = *isActive
	}

	if len(updates) == 0 {
		return 0, nil
	}

	result, err := goqu.Update(TABLE).Prepared(true).Set(updates).Where(
		goqu.C(ISO_CODE).Eq(isoCode),
	).Executor().Exec()

	if err != nil {
		r.logger.Error(
			"COUNTRY_UPDATE_FAILED",
			"there was an issue when updating the country",
			err,
			map[string]any{
				"isoCode": isoCode,
			},
			map[string]any{
				"name":        name,
				"iso3":        iso3,
				"isActive":    isActive,
				"callingCode": callingCode,
			},
		)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		r.logger.Warn(
			"ROWS_AFFECTED_ERROR",
			"there was an issue when trying to fetch number of affected rows post update",
			&err,
			map[string]any{
				"isoCode": isoCode,
			},
			map[string]any{
				"name":        name,
				"iso3":        iso3,
				"isActive":    isActive,
				"callingCode": callingCode,
			},
		)
		return 0, nil
	}

	return rowsAffected, err
}

func (r *Repository) DeleteOne(
	isoCode string,
) error {
	_, err := goqu.Delete(TABLE).Prepared(true).Where(
		goqu.C(ISO_CODE).Eq(isoCode),
	).Executor().Exec()

	return err
}

func (r *Repository) Count(
	search ports.CountriesSearch,
	filters ports.CountriesFilters,
) (int64, error) {

	w := []exp.Expression{}

	if filters.IsActive != nil {
		w = append(w, goqu.C(IS_ACTIVE).Eq(*filters.IsActive))
	}

	searchClauses := r.buildSearchWhereClauses(search)

	if len(searchClauses) > 1 {
		w = append(w, goqu.Or(searchClauses...))
	} else if len(searchClauses) == 1 {
		w = append(w, searchClauses[0])
	}

	count, err := r.goquDB.From(TABLE).Prepared(true).Where(w...).Count()

	if err != nil {
		return 0, err
	}

	return count, nil
}

func (r *Repository) getOrderedExpression(
	sort ports.CountriesSort,
) exp.OrderedExpression {
	var columnName string

	switch sort.Field {
	case ports.COUNTRY_CREATED_AT:
		columnName = CREATED_AT
	case ports.COUNTRY_MODIFIED_AT:
		columnName = MODIFIED_AT
	case ports.COUNTRY_NAME:
		columnName = NAME
	case ports.COUNTRY_ISO:
		columnName = ISO_CODE
	default:
		columnName = ISO_CODE
	}

	if sort.Order == ports.SORT_ASCENDING {
		return goqu.I(columnName).Asc()
	} else {
		return goqu.I(columnName).Desc()
	}

}

func (r *Repository) buildSearchWhereClauses(
	search ports.CountriesSearch,
) []exp.Expression {
	searchExpressions := []exp.Expression{}

	if search.IsoCode != nil && len(strings.TrimSpace(*search.IsoCode)) > 0 {
		searchExpressions = append(
			searchExpressions,
			goqu.C(ISO_CODE).Eq(
				strings.ToUpper(strings.TrimSpace(*search.IsoCode)),
			),
		)
	}

	if search.Name != nil && len(strings.TrimSpace(*search.Name)) > 0 {
		searchExpressions = append(
			searchExpressions,
			goqu.C(NAME).ILike("%"+strings.TrimSpace(*search.Name)+"%"),
		)
	}

	return searchExpressions
}
