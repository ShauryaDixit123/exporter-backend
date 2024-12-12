package quotesrepo

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/pkg/logging"

	"github.com/doug-martin/goqu/v9"
	"github.com/gofrs/uuid"
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

func (r *Repository) InsertRFQ(
	f rdbms.CreateRFQI,
) (string, error) {
	var Id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE_REQUEST_FOR_QUOTE).Rows(
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
			NOTES:                f.Notes,
			CREATED_BY:           f.CreatedBy,
			CREATED_ON:           f.CreatedOn,
		},
	).Returning("id").Executor().ScanStruct(&Id); er != nil {
		return "", er
	}
	return Id.String(), nil
}

func (r *Repository) InsertRFQItems(
	f []rdbms.CreateRFQItemI,
) error {
	if _, er := r.dbClient.Insert(TABLE_REQUEST_FOR_QUOTE_ITEMS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

// re work required in updates
func (r *Repository) UpdateRFQ(
	f rdbms.RFQI,
) error {
	if _, er := r.dbClient.Update(TABLE_REQUEST_FOR_QUOTE).Set(
		goqu.Record{
			TITLE:       f.Title,
			DESCRIPTION: f.Description,
			STATUS:      f.Status,
		}).Where(goqu.Ex{ID: goqu.C(f.ID)}).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) UpdateRequestItem(
	f rdbms.QuoteItemI,
) error {
	if _, er := r.dbClient.Update(TABLE_REQUEST_FOR_QUOTE_ITEMS).Set(
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

func (r *Repository) InsertQuote(
	f rdbms.CreateQuotesI,
) (string, error) {
	var id uuid.UUID
	if _, er := r.dbClient.Insert(TABLE_QUOTES).Rows(f).Returning("id").Executor().ScanStruct(&id); er != nil {
		return "", er
	}
	return id.String(), nil
}

func (r *Repository) InsertQuoteItems(
	f []rdbms.CreateQuotesItemI,
) error {
	if _, er := r.dbClient.Insert(TABLE_QUOTES_ITEMS).Rows(f).Executor().Exec(); er != nil {
		return er
	}
	return nil
}

func (r *Repository) GetRFQsForAccount(
	f rdbms.GetRFQsForAccountI,
) ([]rdbms.RFQI, error) {
	var rfqs []rdbms.RFQI
	if er := r.dbClient.Select(
		goqu.Star(),
	).From(TABLE_REQUEST_FOR_QUOTE).Where(goqu.Ex{ACCOUNT_ID: f.ID}).Executor().ScanStructs(&rfqs); er != nil {
		return nil, er
	}
	return rfqs, nil
}

func (r *Repository) GetRFQ(
	f rdbms.GetRFQI,
) (rdbms.RFQI, error) {
	var rfq rdbms.RFQI
	if _, er := r.dbClient.Select(
		goqu.Star(),
	).From(TABLE_REQUEST_FOR_QUOTE).Where(goqu.Ex{ID: f.ID}).Executor().ScanStruct(&rfq); er != nil {
		return rdbms.RFQI{}, er
	}
	return rfq, nil
}

func (r *Repository) GetRFQItems(
	f rdbms.GetRFQI,
) ([]rdbms.RFQItemI, error) {
	var rfqItems []rdbms.RFQItemI
	if er := r.dbClient.Select(
		goqu.Star(),
	).From(TABLE_REQUEST_FOR_QUOTE_ITEMS).Where(goqu.Ex{RFQ_RFQ_ID: f.ID}).Executor().ScanStructs(&rfqItems); er != nil {
		return nil, er
	}
	return rfqItems, nil
}
