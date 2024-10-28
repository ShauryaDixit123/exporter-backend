package rdbms

import (
	"time"
)

type PurchaseOrder struct {
	Id                   string  `db:"id" json:"id"`
	FlowInstanceId       *string `db:"flow_instance_id" json:"flow_instance_id"`
	FlowInstanceParamsId *string `db:"flow_instance_params_id" json:"flow_instance_params_id"`
	AccountId            int     `db:"account_id" json:"account_id"`
	PONumber             int     `db:"po_number" json:"po_number"`
	DueDate              *string `db:"due_date" json:"due_date"`
	ShipmentTerms        string  `db:"shipment_terms" json:"shipment_terms"`
	TermsAndConditions   string  `db:"terms_and_conditions" json:"terms_and_conditions"`
	Remarks              *string `db:"remarks" json:"remarks"`
	RejectionReason      *string `db:"rejection_reason" json:"rejection_reason"`
	Status               *string `db:"status" json:"status"`
	SupplierId           string  `db:"supplier_id" json:"supplier_id"`
	ShipmentMode         string  `db:"shipment_mode" json:"shipment_mode"`
	PickupUserLocationId int     `db:"pickup_location_id" json:"pickup_location_id"`
	DropUserLocationId   int     `db:"drop_location_id" json:"drop_location_id"`

	BuyerId    string `db:"buyer_id" json:"buyer_id"`
	CreatedBy  string `db:"created_by" json:"created_by"`
	ModifiedBy string `db:"modified_by" json:"modified_by"`
}

type CreatePurchaseOrder struct {
	PurchaseOrder
	WorkflowId *string          `json:"workflow_id"`
	InstanceId *string          `json:"instance_id" db:"instance_id"`
	LineItems  []OrderLineItems `json:"line_items"`
}
type AttachWorkflowI struct {
	AccountId            int     `db:"account_id"`
	FlowInstanceId       *string `db:"flow_instance_id" json:"flow_instance_id"`
	FlowInstanceParamsId *string `db:"flow_instance_params_id" json:"flow_instance_params_id"`
}
type AttachWorkflowReqI struct {
	AccountId        int     `db:"account_id"`
	InstanceId       *string `json:"instance_id" db:"instance_id"`
	WorkflowID       *string `db:"workflow_id" json:"workflow_id"`               // nil that means default workflow
	FlowInstanceType string  `db:"flow_instance_type" json:"flow_instance_type"` // nil means
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
