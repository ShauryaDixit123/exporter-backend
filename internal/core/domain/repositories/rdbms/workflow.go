package rdbms

import (
	"time"

	"github.com/google/uuid"
)

type WorkflowI struct {
	ID        string    `db:"id" json:"id"`
	Name      string    `db:"name" json:"name"`
	Type      string    `db:"type" json:"type"`
	AccountId int       `db:"account_id" json:"account_id"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type CreateWorkflowI struct {
	ID        string        `db:"id" json:"id"`
	Name      string        `db:"name" json:"name"`
	Type      string        `db:"type" json:"type"`
	AccountId int           `db:"account_id" json:"account_id"`
	Flows     []CreateFlowI `json:"flows"`
}

type GetWorkflowI struct {
	ID                  string `db:"id" json:"id"`
	Name                string `db:"name" json:"name"`
	Type                string `db:"type" json:"type"`
	Description         string `db:"description" json:"description"`
	FlowId              int    `db:"flow_id"`
	FlowType            string `db:"type" json:"flow_type"`
	Title               string `db:"title" json:"title"`
	Order               int    `db:"order" json:"order"`
	Active              bool   `db:"active" json:"active"`
	TAT                 int    `db:"tat" json:"tat"`
	FlowParamsId        int    `db:"flow_params_id"`
	FlowParamsName      string `db:"name" json:"flow_params_name"`
	FlowParamsType      string `db:"type" json:"flow_params_type"`
	FlowParamsMandatory bool   `db:"mandatory" json:"flow_params_mandatory"`
}

type FlowI struct {
	ID          string    `db:"id" json:"id"`
	WorkflowID  uuid.UUID `db:"workflow_id" json:"workflow_id"`
	Description string    `db:"description" json:"description"`
	Type        string    `db:"type" json:"type"`
	Title       string    `db:"title" json:"title"`
	Order       int       `db:"order" json:"order"`
	Active      bool      `db:"active" json:"active"`
	TAT         int       `db:"tat" json:"tat"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedBy   string    `db:"updated_by" json:"updated_by"`
	CreatedBy   string    `db:"created_by" json:"created_by"`
}
type FlowAccountsResponseI struct {
	FlowID      string `json:"flow_id" db:"flow_id"`
	WorkflowID  string `json:"workflow_id" db:"workflow_id"`
	Description string `json:"description" db:"description"`
	FlowType    string `json:"flow_type" db:"flow_type"`
	Title       string `json:"title" db:"title"`
	Order       int    `json:"order" db:"order"`
	Active      bool   `json:"active" db:"active"`
	TAT         int    `json:"tat" db:"tat"`
}

type CreateFlowI struct {
	FlowParams []FlowParamI `json:"flow_params"`
	FlowI
}

type CreateWorkflowInstanceI struct {
	Wid       string `db:"workflow_id" json:"wid"`
	AccountId int    `db:"account_id" json:"account_id"`
}
type GetWorkflowByType struct {
	Type      string `db:"type" json:"type"`
	AccountId *int   `db:"account_id" json:"account_id"`
}
type GetFlowInstance struct {
	Type           *string `db:"type" json:"type"`
	InstanceId     *string `db:"instance_id" json:"instance_id"`
	FlowInstanceId *string `db:"flow_instance_id" json:"flow_instance_id"`
}

type GetInstanceAccount struct {
	AccountId int `db:"account_id" json:"account_id"`
	// InstanceId *string `db:"instance_id" json:"instance_id"`
}

type FlowParamI struct {
	// ID        string    `db:"id" json:"id"`
	FlowID    uuid.UUID `db:"flow_id" json:"flow_id"`
	Name      string    `db:"name" json:"name"`
	Type      string    `db:"type" json:"type"`
	Mandatory bool      `db:"mandatory" json:"mandatory"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedBy string    `db:"updated_by" json:"updated_by"`
	CreatedBy string    `db:"created_by" json:"created_by"`
}
type GetFlowParamsResponseI struct {
	FlowParamI
	ID string `db:"id" json:"id"`
}

type FlowInstanceI struct {
	// ID           string    `db:"id" json:"id"`
	WorkflowID   uuid.UUID `db:"workflow_id" json:"workflow_id"`
	Description  string    `db:"description" json:"description"`
	Type         string    `db:"type" json:"type"`
	Title        string    `db:"title" json:"title"`
	Order        int       `db:"order" json:"order"`
	Active       bool      `db:"active" json:"active"`
	TAT          int       `db:"tat" json:"tat"`
	InstanceID   string    `db:"instance_id" json:"instance_id"`
	InstanceType string    `db:"instance_type" json:"instance_type"`
	IsCompleted  bool      `db:"is_completed" json:"is_completed"`
	Status       string    `db:"status" json:"status"`
	AssignedTo   string    `db:"assigned_to" json:"assigned_to"`
	ExpiresAt    time.Time `db:"expires_at" json:"expires_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedBy    string    `db:"updated_by" json:"updated_by"`
	CreatedBy    string    `db:"created_by" json:"created_by"`
}

type UpdateFlowInstanceI struct {
	Id         string     `json:"id" db:"id"`
	Status     *string    `json:"status" db:"status"`
	AssignedTo *string    `json:"assigned_to" db:"assigned_to"`
	Active     *bool      `json:"active" db:"active"`
	ExpiresAt  *time.Time `json:"expires_at" db:"expires_at"`
}
type GetFlowInstanceResponseI struct {
	Id string `json:"id" db:"id"`
	FlowInstanceI
}

type GetFlowInstanceParamsResponseI struct {
	FlowInstanceParamI
	Id string `json:"id" db:"id"`
}
type FlowInstanceParamI struct {
	// ID             string    `db:"id" json:"id"`
	FlowInstanceID string    `db:"flow_instance_id" json:"flow_instance_id"`
	Name           string    `db:"name" json:"name"`
	Type           string    `db:"type" json:"type"`
	Mandatory      bool      `db:"mandatory" json:"mandatory"`
	Value          *string   `db:"value" json:"value"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
	CreatedBy      string    `db:"created_by" json:"created_by"`
}
type UpdateFlowInstanceParamsI struct {
	Id    string `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
}

type GetInstancesI struct {
	AccountId int    `db:"account_id" json:"account_id"`
	Title     string `db:"title" json:"title"`
	Order     int    `db:"order" json:"order"`
}
type FlowInstanceDetails struct {
	FlowInstanceID              string    `db:"flow_instance_id" json:"flow_instance_id"`
	Description                 string    `db:"description" json:"description"`
	FlowInstanceType            string    `db:"flow_instance_type" json:"flow_instance_type"`
	FlowInstanceParamsType      string    `db:"flow_instance_params_type" json:"flow_instance_params_type"`
	Title                       string    `db:"title" json:"title"`
	Order                       int       `db:"order" json:"order"`
	Active                      bool      `db:"active" json:"active"`
	TAT                         int       `db:"tat" json:"tat"`
	FlowInstanceStatus          string    `db:"status" json:"status"`
	IsCompleted                 bool      `db:"is_completed" json:"is_completed"`
	AssignedTo                  string    `db:"assigned_to" json:"assigned_to"`
	WorkflowID                  string    `db:"workflow_id" json:"workflow_id"`
	FlowInstanceParamsID        string    `db:"flow_instance_params_id" json:"flow_instance_params_id"`
	FlowInstanceParamsName      string    `db:"flow_instance_params_name" json:"flow_instance_params_name"`
	FlowInstanceParamsValue     *string   `db:"flow_instance_params_value" json:"flow_instance_params_value"`
	FlowInstanceParamsMandatory bool      `db:"flow_instance_params_mandatory" json:"flow_instance_params_mandatory"`
	CreatedAt                   time.Time `db:"created_at" json:"created_at"` // Example additional fields if needed
	UpdatedAt                   time.Time `db:"updated_at" json:"updated_at"` // Example additional fields if needed
}
type GetFlowsForAccountI struct {
	AccountId int  `db:"account_id" json:"account_id"`
	PreOrder  bool `db:"pre_order" json:"pre_order"`
}

type GroupedFlowInstancesResponse struct {
	ID                 string                           `json:"flow_instance_id"`
	Description        string                           `json:"description"`
	Type               string                           `json:"flow_instance_type"`
	FlowInstanceParams []GetFlowInstanceParamsResponseI `json:"flow_instance_params"`
}
