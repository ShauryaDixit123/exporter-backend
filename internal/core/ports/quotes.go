package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsQuotesRepository interface {
	InsertRequest(
		f rdbms.CreateQuotesI,
	) (int, error)
	InsertRequestItems(
		f []rdbms.CreateQuotesItemI,
	) error
	UpdateRequest(
		f rdbms.QoutesI,
	) error
	UpdateRequestItem(
		f rdbms.QuoteItemI,
	) error
}
