package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/domain/services"
)

type RdbmsCountriesRepository interface {
	SelectOne(isoCode string) (rdbms.Country, bool, error)
	SelectMany(
		limit uint,
		offset uint,
		search CountriesSearch,
		sort CountriesSort,
		filters CountriesFilters,
	) ([]rdbms.Country, error)
	InsertOne(
		isoCode string,
		name string,
		iso3 *string,
		callingCode uint,
		isActive bool,
	) error
	UpdateOne(
		isoCode string,
		name *string,
		iso3 *string,
		callingCode *uint,
		isActive *bool,
	) (int64, error)
	DeleteOne(
		isoCode string,
	) error
	Count(
		search CountriesSearch,
		filters CountriesFilters,
	) (int64, error)
}

type CountriesService interface {
	Fetch(isoCode string) (services.Country, bool, error)
	FetchMany(
		pageNumber uint,
		itemsPerPage uint,
		searchTerm *string,
		sort CountriesSort,
		filters CountriesFilters,
	) (services.Countries, error)
	Create(
		isoCode string,
		name string,
		iso3 *string,
		callingCode uint,
		symbol string,
		isActive bool,
	) (services.Country, error)
	Modify(
		isoCode string,
		name *string,
		iso3 *string,
		callingCode *uint,
		isActive *bool,
	) (services.Country, error)
	Remove(
		isoCode string,
	) error
}

type CountriesSortField int

const (
	COUNTRY_CREATED_AT  CountriesSortField = iota
	COUNTRY_MODIFIED_AT CountriesSortField = iota
	COUNTRY_NAME        CountriesSortField = iota
	COUNTRY_ISO         CountriesSortField = iota
)

type CountriesSort struct {
	Field CountriesSortField
	Order SortOrder
}

type CountriesFilters struct {
	IsActive *bool
}

type CountriesSearch struct {
	IsoCode *string
	Name    *string
}
