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
	Type       string `db:"type" json:"type"`
	InstanceId string `db:"instance_id" json:"instance_id"`
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
type GetFlowInstanceResponseI struct {
	Id string `json:"id" db:"id"`
	FlowInstanceI
}

type GetFlowInstanceParamsResponseI struct {
	Id string `json:"id" db:"id"`
	FlowInstanceParamI
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
