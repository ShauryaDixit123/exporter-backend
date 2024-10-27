package rdbms

import (
	"time"

	"github.com/gofrs/uuid"
)

type AccountI struct {
	Id            int       `db:"id" json:"id"`
	PrimaryUserID uuid.UUID `db:"primary_user_id" json:"primary_user_id"`
	IsActive      bool      `db:"is_active" json:"is_active"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	ModifiedAt    time.Time `db:"modified_at" json:"modified_at"`
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
	AccountId int    `json:"account_id" db:"account_id"`
	UserId    string `json:"user_id" db:"user_id"`
}

type CreateFlowInstanceAccountI struct {
	FlowInstanceId string `db:"flow_instance_id" json:"flow_instance_id"`
	AccountId      int    `db:"account_id" json:"account_id"`
}
