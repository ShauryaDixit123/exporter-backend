package quotesrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

	"github.com/doug-martin/goqu/v9"
)

type Repository struct {
	logger   logging.Logger
	dbClient *goqu.Database
}

func New(
	logger logging.Logger,
	dbClient *goqu.Database,
) *Repository {
	return &Repository{
		logger:   logger,
		dbClient: dbClient,
	}
}

func (r *Repository) Insert(
	f rdbms.CreateQuotesI,
) (int, error) {
	var Id int
	if _, er := r.dbClient.Insert(TABLE).Rows(
		goqu.Record{
			ACCOUNT_ID:           f.AccountID,
			BUYER_ID:             f.BuyerID,
			SUPPLIER_ID:          f.SupplierID,
			TITLE:                f.Title,
			DESCRIPTION:          f.Description,
			INCO_TERMS:           f.IncoTerms,
			PICKUP_LOCATION_ID:   f.PickupLocationID,
			DROP_LOCATION_ID:     f.DropLocationID,
			PAYMENT_TERMS:        f.PaymentTerms,
			ACTIVE:               f.Active,
			TAT:                  f.TAT,
			DUE_DATE:             f.DueDate,
			STATUS:               f.Status,
			TERMS_AND_CONDITIONS: f.TermsAndConditions,
			CREATED_BY:           f.CreatedBy,
			CREATED_ON:           f.CreatedOn,
		},
	).Returning("id").Executor().ScanStruct(&Id); er != nil {
		return 0, er
	}
	return Id, nil
}

func (r *Repository) InsertItems(
	f []rdbms.CreateQuotesItemI,
) error {
	if _, er := r.dbClient.Insert(TABLE_ITEMS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

// re work required in updates
func (r *Repository) Update(
	f rdbms.QoutesI,
) error {
	if _, er := r.dbClient.Update(TABLE).Set(
		goqu.Record{
			TITLE:       f.Title,
			DESCRIPTION: f.Description,
			STATUS:      f.Status,
		}).Where(goqu.Ex{ID: goqu.C(f.ID)}).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) UpdateItem(
	f rdbms.QuoteItemI,
) error {
	if _, er := r.dbClient.Update(TABLE_ITEMS).Set(
		goqu.Record{
			QUANTITY:    f.Quantity,
			TITLE:       f.Title,
			DESCRIPTION: f.Description,
			RATE:        f.Rate,
			RATE_UNIT:   f.RateUnit,
		},
	).Where(goqu.Ex{ID: f.ID}).Executor().Exec(); er != nil {
		return er
	}
	return nil
}