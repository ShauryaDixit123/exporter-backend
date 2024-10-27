package currenciessrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/domain/services"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"
)

type Service struct {
	logger               logging.Logger
	currenciesRepository ports.RdbmsCurrenciesRepository
}

func New(
	logger logging.Logger,
	currenciesRepository ports.RdbmsCurrenciesRepository,
) *Service {
	return &Service{
		logger:               logger,
		currenciesRepository: currenciesRepository,
	}
}

func (s *Service) Fetch(isoCode string) (services.Currency, bool, error) {
	return s.fetchOneCurrency(isoCode)
}

func (s *Service) FetchMany(
	pageNumber uint,
	itemsPerPage uint,
	searchTerm *string,
	sort *ports.CurrenciesSort,
	filters *ports.CurrenciesFilters,
) ([]services.Currency, error) {
	if pageNumber == 0 {
		pageNumber = 1
	}

	offset := (itemsPerPage * (pageNumber - 1))

	sortBy := ports.CurrenciesSort{
		Field: ports.CURRENCY_NAME,
		Order: ports.SORT_ASCENDING,
	}

	if sort != nil {
		sortBy = *sort
	}

	search := ports.CurrenciesSearch{
		IsoCode: searchTerm,
		Name:    searchTerm,
	}

	dbFilters := ports.CurrenciesFilters{}

	if filters != nil {
		dbFilters = *filters
	}

	repoCurrencies, err := s.currenciesRepository.SelectMany(
		itemsPerPage,
		offset,
		search,
		sortBy,
		dbFilters,
	)

	if err != nil {
		return []services.Currency{}, err
	}

	currencies := make([]services.Currency, 0)

	for _, c := range repoCurrencies {
		currencies = append(currencies, s.mapRepoDomainToService(c))
	}

	return currencies, nil

}

func (s *Service) Create(
	isoCode string,
	name string,
	symbol string,
	isActive bool,
) (services.Currency, error) {
	err := s.currenciesRepository.InsertOne(
		isoCode,
		name,
		symbol,
		isActive,
	)

	if err != nil {
		return services.Currency{}, err
	}

	currency, exists, err := s.fetchOneCurrency(isoCode)

	if err != nil {
		return services.Currency{}, err
	}

	if !exists {
		return services.Currency{}, fmt.Errorf(
			"could not find the currency just created with iso code: %s",
			isoCode,
		)
	}

	return currency, nil

}

func (s *Service) Modify(
	isoCode string,
	name *string,
	symbol *string,
	isActive *bool,
) (services.Currency, error) {
	_, err := s.currenciesRepository.UpdateOne(
		isoCode,
		name,
		symbol,
		isActive,
	)

	if err != nil {
		return services.Currency{}, err
	}

	currency, exists, err := s.fetchOneCurrency(isoCode)

	if err != nil {
		return services.Currency{}, err
	}

	if !exists {
		return services.Currency{}, fmt.Errorf(
			"could not find the currency just updated with iso code: %s",
			isoCode,
		)
	}

	return currency, nil
}

func (s *Service) Remove(
	isoCode string,
) error {
	err := s.currenciesRepository.DeleteOne(isoCode)

	return err
}

func (s *Service) mapRepoDomainToService(
	c rdbms.Currency,
) services.Currency {
	return services.Currency{
		IsoCode:  c.IsoCode,
		Name:     c.Name,
		Symbol:   c.Symbol,
		IsActive: c.IsActive,
	}
}

func (s *Service) fetchOneCurrency(isoCode string) (services.Currency, bool, error) {
	repoCurrency, exists, err := s.currenciesRepository.SelectOne(isoCode)

	if err != nil {
		return services.Currency{}, false, err
	}

	if !exists {
		return services.Currency{}, false, nil
	}

	return s.mapRepoDomainToService(repoCurrency), true, nil
}
