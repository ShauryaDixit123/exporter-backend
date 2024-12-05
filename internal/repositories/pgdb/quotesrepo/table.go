package quotesrepo

const (
	ID                   = "id"
	ACCOUNT_ID           = "account_id"
	BUYER_ID             = "buyer_id"
	SUPPLIER_ID          = "supplier_id"
	TITLE                = "title"
	DESCRIPTION          = "description"
	INCO_TERMS           = "inco_terms"
	PICKUP_LOCATION_ID   = "pickup_location_id"
	DROP_LOCATION_ID     = "drop_location_id"
	PAYMENT_TERMS        = "payment_terms"
	ACTIVE               = "active"
	TAT                  = "tat"
	DUE_DATE             = "due_date"
	STATUS               = "status"
	TERMS_AND_CONDITIONS = "terms_and_conditions"
	CREATED_BY           = "created_by"
	CREATED_ON           = "created_on"
	MODIFIED_AT          = "modified_at"
)

const TABLE_REQUEST_FOR_QUOTE = "request_for_quote"
const TABLE_REQUEST_FOR_QUOTE_ITEMS = "request_for_quote_items"
const TABLE_QUOTES = "quotes"
const TABLE_QUOTES_ITEMS = "quotes_items"

const (
	// ID            = "ID"
	QUOTE_ID  = "QUOTE_ID"
	ITEM_CODE = "ITEM_CODE"
	STORE_ID  = "STORE_ID"
	// TITLE         = "TITLE"
	// DESCRIPTION   = "DESCRIPTION"
	QUANTITY      = "QUANTITY"
	QUANTITY_UNIT = "QUANTITY_UNIT"
	RATE          = "RATE"
	RATE_UNIT     = "RATE_UNIT"
	DELIVERY_DATE = "DELIVERY_DATE"
)

const (
	RFQ_RFQ_ID               = "rfq_id"
	RFQ_SUPPLIER_ID          = "supplier_id"
	RFQ_ACTIVE               = "active"
	RFQ_DUE_DATE             = "due_date"
	RFQ_STATUS               = "status"
	RFQ_TERMS_AND_CONDITIONS = "terms_and_conditions"
	RFQ_REMARKS              = "remarks"
	RFQ_REJECTION_REASON     = "rejection_reason"
	RFQ_CREATED_ON           = "created_on"
	RFQ_MODIFIED_AT          = "modified_at"
)
const (
	QUOTE_QUOTE_ID      = "quote_id"
	QUOTE_RFQ_ITEM_ID   = "rfq_item_id"
	QUOTE_RATE          = "rate"
	QUOTE_RATE_UNIT     = "rate_unit"
	QUOTE_IMAGE_ID      = "image_id"
	QUOTE_QUANTITY      = "quantity"
	QUOTE_QUANTITY_UNIT = "quantity_unit"
	QUOTE_DELIVERY_DATE = "delivery_date"
	QUOTE_STATUS        = "status"
	QUOTE_CREATED_ON    = "created_on"
	QUOTE_MODIFIED_AT   = "modified_at"
)
