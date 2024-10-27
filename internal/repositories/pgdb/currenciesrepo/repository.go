package currenciesrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"
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
) (rdbms.Currency, bool, error) {
	var currency rdbms.Currency

	exists, err := r.goquDB.From(TABLE).Prepared(true).Select(
		ISO_CODE,
		NAME,
		SYMBOL,
		IS_ACTIVE,
		CREATED_AT,
		MODIFIED_AT,
	).Where(
		goqu.C(ISO_CODE).Eq(isoCode),
	).ScanStruct(&currency)

	if err != nil {
		return rdbms.Currency{}, false, err
	}

	if !exists {
		return rdbms.Currency{}, false, nil
	}

	return currency, true, nil
}

func (r *Repository) SelectMany(
	limit uint,
	offset uint,
	search ports.CurrenciesSearch,
	sort ports.CurrenciesSort,
	filters ports.CurrenciesFilters,
) ([]rdbms.Currency, error) {
	var currencies []rdbms.Currency

	w := []exp.Expression{}

	if filters.IsActive != nil {
		w = append(w, goqu.I(IS_ACTIVE).Eq(*filters.IsActive))
	}

	searchClauses := r.buildSearchWhereClauses(search)

	if len(searchClauses) > 1 {
		w = append(w, goqu.Or(searchClauses...))
	} else if len(searchClauses) == 1 {
		w = append(w, searchClauses[0])
	}

	query := r.goquDB.From(TABLE).Prepared(true).Select(
		ISO_CODE,
		NAME,
		SYMBOL,
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

	err := query.ScanStructs(&currencies)

	if err != nil {
		querySQL, params, qerr := query.ToSQL()

		r.logger.Error(
			"CURRENCY_LIST_SELECTION_FAILED",
			"there was an issue when updating the country",
			err,
			map[string]any{},
			map[string]any{
				"query":  querySQL,
				"params": params,
				"error":  qerr.Error(),
			},
		)

		return []rdbms.Currency{}, err
	}

	return currencies, nil
}

func (r *Repository) InsertOne(
	isoCode string,
	name string,
	symbol string,
	isActive bool,
) error {
	_, err := r.goquDB.Insert(TABLE).Prepared(true).Rows(
		goqu.Record{
			ISO_CODE:  isoCode,
			NAME:      name,
			SYMBOL:    symbol,
			IS_ACTIVE: isActive,
		},
	).Executor().Exec()

	return err
}
func (r *Repository) UpdateOne(
	isoCode string,
	name *string,
	symbol *string,
	isActive *bool,
) (int64, error) {
	updates := goqu.Record{}

	if name != nil {
		updates[NAME] = *name
	}

	if symbol != nil {
		updates[SYMBOL] = *symbol
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
			"CURRENCY_UPDATE_FAILED",
			"there was an issue when updating the country",
			err,
			map[string]any{
				"isoCode": isoCode,
			},
			map[string]any{
				"name":     name,
				"symbol":   symbol,
				"isActive": isActive,
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
				"name":     name,
				"symbol":   symbol,
				"isActive": isActive,
			},
		)
		return 0, nil
	}

	return rowsAffected, err
}

func (r *Repository) DeleteOne(
	isoCode string,
) error {
	_, err := r.goquDB.Delete(TABLE).Where(
		goqu.C(ISO_CODE).Eq(isoCode),
	).Executor().Exec()

	if err != nil {
		r.logger.Error(
			"CURRENCY_DELETION_FAILED",
			"there was an issue when deleting the currency",
			err,
			map[string]any{
				"isoCode": isoCode,
			},
			map[string]any{},
		)
	}

	return err
}

func (r *Repository) getOrderedExpression(
	sort ports.CurrenciesSort,
) exp.OrderedExpression {
	var columnName string

	switch sort.Field {
	case ports.CURRENCY_CREATED_AT:
		columnName = CREATED_AT
	case ports.CURRENCY_MODIFIED_AT:
		columnName = MODIFIED_AT
	case ports.CURRENCY_NAME:
		columnName = NAME
	case ports.CURRENCY_ISO:
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
	search ports.CurrenciesSearch,
) []exp.Expression {
	searchExpressions := []exp.Expression{}

	if search.IsoCode != nil && len(strings.TrimSpace(*search.IsoCode)) > 0 {
		searchExpressions = append(
			searchExpressions,
			goqu.L(
				fmt.Sprintf("%s @@ TO_TSQUERY(?)", ISO_CODE),
				strings.TrimSpace(*search.IsoCode),
			),
		)
	}

	if search.Name != nil && len(strings.TrimSpace(*search.Name)) > 0 {
		searchExpressions = append(
			searchExpressions,
			goqu.L(
				fmt.Sprintf("%s @@ TO_TSQUERY(?)", NAME),
				strings.TrimSpace(*search.Name),
			),
		)
	}

	return searchExpressions
}
