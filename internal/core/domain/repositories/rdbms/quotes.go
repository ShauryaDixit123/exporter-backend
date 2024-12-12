package rdbms

import (
	"time"
)

type CreateQuotesI struct {
	RFQID              string    `json:"rfq_id" db:"rfq_id"`
	SupplierID         string    `json:"supplier_id" db:"supplier_id"`
	Active             bool      `json:"active" db:"active"`
	DueDate            time.Time `json:"due_date" db:"due_date"`
	Status             string    `json:"status" db:"status"` // Assuming 18 char max
	TermsAndConditions string    `json:"terms_and_conditions" db:"terms_and_conditions"`
	Remarks            string    `json:"remarks" db:"remarks"`
	RejectionReason    string    `json:"rejection_reason" db:"rejection_reason"`
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
type CreateRFQItemI struct {
	QuoteID      string    `json:"quote_id" db:"quote_id"`
	RFQID        string    `json:"rfq_id" db:"rfq_id"`
	Rate         int       `json:"rate" db:"rate"`
	RateUnit     string    `json:"rate_unit" db:"rate_unit"` // Assuming 10 char max
	ImageID      *string   `json:"image_id" db:"image_id"`   // Nullable
	Quantity     int       `json:"quantity" db:"quantity"`
	QuantityUnit string    `json:"quantity_unit" db:"quantity_unit"` // Assuming 16 char max
	DeliveryDate string    `json:"delivery_date" db:"delivery_date"` // Assuming 16 char max
	Status       string    `json:"status" db:"status"`               // Assuming 18 char max
	CreatedOn    time.Time `json:"created_on" db:"created_on"`
	ModifiedAt   time.Time `json:"modified_at" db:"modified_at"`
}
type RFQItemI struct {
	ID string `json:"id" db:"id"`
	CreateRFQItemI
}

type CreateRFQI struct {
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
	Notes              string    `json:"notes" db:"notes"`
	CreatedOn          time.Time `json:"created_on" db:"created_on"`
	ModifiedAt         time.Time `json:"modified_at" db:"modified_at"`
}
type CreateRFQRequestI struct {
	CreateRFQI
	Items []CreateRFQItemI
}
type RFQI struct {
	ID string `json:"id" db:"id"`
	CreateRFQI
}
type CreateQuoteRequestI struct {
	CreateQuotesI
	Items []CreateQuotesItemI `json:"items"`
}
type RFQResponseI struct {
	RFQI
	Items []RFQItemI `json:"items"`
}

type GetRFQsForAccountI struct {
	ID int `json:"id" db:"id"`
}
type GetRFQI struct {
	ID string `json:"id" db:"id"`
}
