package ports

import "exporterbackend/internal/core/domain/repositories/rdbms"

type RdbmsQuotesRepository interface {
	Insert(
		f rdbms.CreateQuotesI,
	) (int, error)
	InsertItems(
		f rdbms.CreateQuotesItemI,
	) error
	Update(
		f rdbms.QoutesI,
	) error
	UpdateItems(
		f rdbms.QuoteItemI,
	) error
}
