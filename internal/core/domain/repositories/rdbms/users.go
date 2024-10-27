package rdbms

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type CreateUserI struct {
	Name              string    `db:"name" json:"name"`
	Email             string    `db:"email" json:"email"`
	Password          string    `db:"password" json:"password"`
	IsParent          bool      `db:"is_parent" json:"is_parent"`
	RoleId            int       `db:"role_id"`
	Role              string    `json:"role"`
	PrimaryLocationID uuid.UUID `db:"primary_location_id" json:"primary_location_id"`
	CreatedBy         string
}

type UserI struct {
	Id                uuid.UUID `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	Email             string    `db:"email" json:"email"`
	Password          string    `db:"password" json:"-"`
	IsParent          bool      `db:"is_parent" json:"is_parent"`
	PrimaryLocationID uuid.UUID `db:"primary_location_id" json:"primary_location_id"`
	AccessToken       uuid.UUID `db:"access_token" json:"access_token"`
	RoleId            int       `db:"role_id"`
	Role              string    `json:"role"`
	IsActive          bool      `db:"is_active" json:"is_active"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	ModifiedAt        time.Time `db:"modified_at" json:"modified_at"`
}

type RoleI struct {
	Id         int       `json:"id" db:"id"`
	Role       string    `json:"role" db:"role"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`
}

type Id struct {
	Id string `db:"id"`
}
type IdInt struct {
	Id int `db:"id"`
}

type GetAccountUsersI struct {
	AccountId int `db:"account_id" json:"account_id"`
}
