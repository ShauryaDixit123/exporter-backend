package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsQuotesRepository interface {
	InsertRFQ(
		f rdbms.CreateRFQI,
	) (int, error)
	InsertRFQItems(
		f []rdbms.CreateRFQItemI,
	) error
	UpdateRFQ(
		f rdbms.RFQI,
	) error
	UpdateRequestItem(
		f rdbms.QuoteItemI,
	) error
}
