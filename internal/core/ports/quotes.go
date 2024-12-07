package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsQuotesRepository interface {
	InsertRFQ(
		f rdbms.CreateRFQI,
	) (string, error)
	InsertRFQItems(
		f []rdbms.CreateRFQItemI,
	) error
	UpdateRFQ(
		f rdbms.RFQI,
	) error
	UpdateRequestItem(
		f rdbms.QuoteItemI,
	) error
	InsertQuote(
		f rdbms.CreateQuotesI,
	) (string, error)
	InsertQuoteItems(
		f []rdbms.CreateQuotesItemI,
	) error
	GetRFQsForAccount(
		f rdbms.GetRFQsForAccountI,
	) ([]rdbms.RFQI, error)
	GetRFQ(
		f rdbms.GetRFQI,
	) (rdbms.RFQI, error)
	GetRFQItems(
		f rdbms.GetRFQI,
	) ([]rdbms.RFQItemI, error)
}
type QuotesService interface {
	CreateRFQ(
		f rdbms.CreateRFQRequestI,
	) error
	CreateQuote(
		f rdbms.CreateQuoteRequestI,
	) error
	GetRfQsForAccount(
		f rdbms.GetRFQsForAccountI,
	) ([]rdbms.RFQI, error)
	GetRFQ(
		f rdbms.GetRFQI,
	) (rdbms.RFQResponseI, error)
}
