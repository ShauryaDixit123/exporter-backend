package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/domain/services"
)

type RdbmsCurrenciesRepository interface {
	SelectOne(isoCode string) (rdbms.Currency, bool, error)
	SelectMany(
		limit uint,
		offset uint,
		search CurrenciesSearch,
		sort CurrenciesSort,
		filters CurrenciesFilters,
	) ([]rdbms.Currency, error)
	InsertOne(
		isoCode string,
		name string,
		symbol string,
		isActive bool,
	) error
	UpdateOne(
		isoCode string,
		name *string,
		symbol *string,
		isActive *bool,
	) (int64, error)
	DeleteOne(
		isoCode string,
	) error
}

type CurrenciesService interface {
	Fetch(isoCode string) (services.Currency, bool, error)
	FetchMany(
		pageNumber uint,
		itemsPerPage uint,
		searchTerm *string,
		sort *CurrenciesSort,
		filters *CurrenciesFilters,
	) ([]services.Currency, error)
	Create(
		isoCode string,
		name string,
		symbol string,
		isActive bool,
	) (services.Currency, error)
	Modify(
		isoCode string,
		name *string,
		symbol *string,
		isActive *bool,
	) (services.Currency, error)
	Remove(
		isoCode string,
	) error
}

type CurrenciesSortField int

const (
	CURRENCY_CREATED_AT  CurrenciesSortField = iota
	CURRENCY_MODIFIED_AT CurrenciesSortField = iota
	CURRENCY_NAME        CurrenciesSortField = iota
	CURRENCY_ISO         CurrenciesSortField = iota
)

type CurrenciesSort struct {
	Field CurrenciesSortField
	Order SortOrder
}

type CurrenciesFilters struct {
	IsActive *bool
}

type CurrenciesSearch struct {
	IsoCode *string
	Name    *string
}
