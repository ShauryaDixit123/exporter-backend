package quotessrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
)

type Service struct {
	logger     logging.Logger
	quotesRepo ports.RdbmsQuotesRepository
}

func New(logger logging.Logger,
	quotesRepo ports.RdbmsQuotesRepository,
) *Service {
	return &Service{
		logger:     logger,
		quotesRepo: quotesRepo,
	}
}

func (s *Service) CreateRFQ(
	f rdbms.CreateRFQRequestI,
) error {
	id, er := s.quotesRepo.InsertRFQ(rdbms.CreateRFQI{
		AccountID:          f.AccountID,
		BuyerID:            f.BuyerID,
		Title:              f.Title,
		Description:        f.Description,
		IncoTerms:          f.IncoTerms,
		PickupLocationID:   f.PickupLocationID,
		DropLocationID:     f.DropLocationID,
		PaymentTerms:       f.PaymentTerms,
		Active:             f.Active,
		TAT:                f.TAT,
		DueDate:            f.DueDate,
		Status:             "pending",
		TermsAndConditions: f.TermsAndConditions,
		CreatedBy:          f.CreatedBy,
		CreatedOn:          f.CreatedOn,
	})
	if er != nil {
		return er
	}
	for i := range f.Items {
		f.Items[i].RFQID = id
	}
	if er := s.quotesRepo.InsertRFQItems(f.Items); er != nil {
		return er
	}
	return nil
}

func (s *Service) CreateQuote(
	f rdbms.CreateQuoteRequestI,
) error {
	id, er := s.quotesRepo.InsertQuote(rdbms.CreateQuotesI{
		RFQID:              f.RFQID,
		SupplierID:         f.SupplierID,
		Status:             "pending",
		TermsAndConditions: f.TermsAndConditions,
		Remarks:            f.Remarks,
		RejectionReason:    f.RejectionReason,
	})
	if er != nil {
		return er
	}
	for i := range f.Items {
		f.Items[i].QuoteID = id
	}
	if er != nil {
		return er
	}

	if er := s.quotesRepo.InsertQuoteItems(f.Items); er != nil {
		return er
	}
	return nil
}

func (s *Service) GetRfQsForAccount(
	f rdbms.GetRFQsForAccountI,
) ([]rdbms.RFQI, error) {
	return s.quotesRepo.GetRFQsForAccount(f)
}

func (s *Service) GetRFQ(
	f rdbms.GetRFQI,
) (rdbms.RFQResponseI, error) {
	var resp rdbms.RFQResponseI
	res, er := s.quotesRepo.GetRFQ(f)
	if er != nil {
		return rdbms.RFQResponseI{}, er
	}

	items, er := s.quotesRepo.GetRFQItems(f)
	if er != nil {
		return rdbms.RFQResponseI{}, er
	}
	resp.RFQI = res
	resp.Items = items
	return resp, nil
}
