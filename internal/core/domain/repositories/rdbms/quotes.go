package rdbms

import "time"

type CreateQuotesI struct {
	AccountID          int       `json:"account_id" db:"account_id"`
	BuyerID            string    `json:"buyer_id" db:"buyer_id"`
	SupplierID         string    `json:"supplier_id" db:"supplier_id"`
	Title              string    `json:"title" db:"title"`
	Description        string    `json:"description" db:"description"`
	IncoTerms          string    `json:"inco_terms" db:"inco_terms"`
	PickupLocationID   int       `json:"pickup_location_id" db:"pickup_location_id"`
	DropLocationID     int       `json:"drop_location_id" db:"drop_location_id"`
	PaymentTerms       string    `json:"payment_terms" db:"payment_terms"`
	Active             bool      `json:"active" db:"active"`
	TAT                int       `json:"tat" db:"tat"` // Turnaround time
	DueDate            time.Time `json:"due_date" db:"due_date"`
	Status             string    `json:"status" db:"status"`
	TermsAndConditions string    `json:"terms_and_conditions" db:"terms_and_conditions"`
	CreatedBy          string    `json:"created_by" db:"created_by"`
	CreatedOn          time.Time `json:"created_on" db:"created_on"`
	ModifiedAt         time.Time `json:"modified_at" db:"modified_at"`
}
type QoutesI struct {
	ID string `json:"id" db:"id"`
	CreateQuotesI
}

type CreateQuotesItemI struct {
	QuoteID      string    `json:"quote_id" db:"quote_id"`
	ItemCode     string    `json:"item_code" db:"item_code"`
	StoreID      string    `json:"store_id" db:"store_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Quantity     int       `json:"quantity" db:"quantity"`
	QuantityUnit string    `json:"quantity_unit" db:"quantity_unit"`
	Rate         int       `json:"rate" db:"rate"`
	RateUnit     string    `json:"rate_unit" db:"rate_unit"`
	DeliveryDate string    `json:"delivery_date" db:"delivery_date"`
	CreatedOn    time.Time `json:"created_on" db:"created_on"`
	ModifiedAt   time.Time `json:"modified_at" db:"modified_at"`
}
type QuoteItemI struct {
	ID string `json:"id" db:"id"`
	CreateQuotesItemI
}