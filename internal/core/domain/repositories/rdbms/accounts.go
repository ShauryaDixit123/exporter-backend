package rdbms

import (
	"time"

	"github.com/google/uuid"
)

type AccountI struct {
	Id                       int       `db:"id" json:"id"`
	PrimaryUserID            string    `db:"primary_user_id" json:"primary_user_id"`
	IsActive                 bool      `db:"is_active" json:"is_active"`
	GstNo                    int       `db:"gst_no" json:"gst_no"`
	DefaultWorkflowPreOrder  string    `db:"default_workflow_pre_order" json:"default_workflow_pre_order"`
	DefaultWorkflowPostOrder string    `db:"default_workflow_post_order" json:"default_workflow_post_order"`
	CreatedAt                time.Time `db:"created_at" json:"created_at"`
	ModifiedAt               time.Time `db:"modified_at" json:"modified_at"`
}
type AccountsUsersMap struct {
	ID         int       `db:"id" json:"id"`
	UserID     uuid.UUID `db:"user_id" json:"user_id"`
	AccountID  int       `db:"account_id" json:"account_id"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`
}
type CreateAccountI struct {
	PrimaryUserID uuid.UUID `db:"primary_user_id" json:"primary_user_id"`
	IsActive      bool      `db:"is_active" json:"is_active"`
}
type CreateAccountUserI struct {
	AccountId int       `json:"account_id" db:"account_id"`
	UserId    uuid.UUID `json:"user_id" db:"user_id"`
}

type CreateFlowInstanceAccountI struct {
	FlowInstanceId string `db:"instance_id" json:"instance_id"`
	AccountId      int    `db:"account_id" json:"account_id"`
}

type GetUserForAccountReq struct {
	Role      string `db:"role" json:"role"`
	AccountId int    `db:"account_id" json:"account_id"`
}

type GetUserForAccount struct {
	RoleId    int `db:"role_id" json:"role_id"`
	AccountId int `db:"account_id" json:"account_id"`
}
