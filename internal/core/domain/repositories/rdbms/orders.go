package rdbms

import "time"

type PurchaseOrder struct {
	Id                   string  `db:"id"`
	UserId               int     `db:"user_id"`
	FlowInstanceId       string  `db:"flow_instance_id" json:"flow_instance_id"`
	FlowInstanceParamsId string  `db:"flow_instance_params_id" json:"flow_instance_params_id"`
	AccountId            int     `db:"account_id"`
	PONumber             string  `db:"po_number"`
	DueDate              *string `db:"due_date"`
	ShipmentTerms        string  `db:"shipment_terms"`
	TermsAndConditions   string  `db:"terms_and_conditions"`
	Remarks              *string `db:"remarks"`
	RejectionReason      *string `db:"rejection_reason"`
	Status               *string `db:"status"`
	SupplierId           int     `db:"supplier_id"`
	ShipmentMode         string  `db:"shipment_mode"`
	Pol                  string  `db:"pickup_location_id"`
	Pod                  string  `db:"drop_location_id"`
	CreatedBy            string  `db:"created_by"`
	ModifiedBy           string  `db:"modified_by"`
}

type OrderLineItems struct {
	Id           string `db:"id"`
	PoId         string `db:"po_id"`
	SoId         string `db:"so_id"`
	LiRefId      string `db:"li_ref_id"`
	ItemCode     string `db:"item_code"`
	Description  string `db:"description"`
	BatchCount   *int   `db:"batch_count"`
	Quantity     int    `db:"quantity"`
	DeliveryDate string `db:"delivery_date"`
	Status       string `db:"status"`
	OG           bool   `db:"og"`
	CreatedOn    string `db:"created_on"`
	ModifedAt    string `db:"modified_at"`
}

type SalesOrder struct {
	Id                   string  `db:"id"`
	AccountId            int     `db:"account_id"`
	POId                 string  `db:"po_id"`
	SONumber             string  `db:"so_number"`
	SupplierId           int     `db:"supplier_id"`
	DueDate              string  `db:"due_date"`
	FlowInstanceId       string  `db:"flow_instance_id" json:"flow_instance_id"`
	FlowInstanceParamsId string  `db:"flow_instance_params_id" json:"flow_instance_params_id"`
	Status               *string `db:"status"`
	// RejectionReason    *string                   `db:"rejection_reason"`
	ModifiedBy int       `db:"modified_by"`
	CreatedBy  int       `db:"created_by"`
	CreatedOn  time.Time `db:"created_on"`
}
