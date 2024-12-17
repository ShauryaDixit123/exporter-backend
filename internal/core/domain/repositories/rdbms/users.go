package rdbms

import (
	"time"
)

type CreateUserI struct {
	Name              string `db:"name" json:"name"`
	Email             string `db:"email" json:"email"`
	Password          string `db:"password" json:"password"`
	IsParent          bool   `db:"is_parent" json:"is_parent"`
	RoleId            int    `db:"role_id"`
	Role              string `json:"role"`
	PrimaryLocationID string `db:"primary_location_id" json:"primary_location_id"`
	CreatedBy         string `json:"created_by"db:"created_by"`
}

type CreateUserRequestI struct {
	IsInvited bool `json:"is_invited"`
	CreateUserI
}

type UserI struct {
	Id                string    `db:"id" json:"id"`
	Name              string    `db:"name" json:"name"`
	Email             string    `db:"email" json:"email"`
	Password          string    `db:"password" json:"-"`
	IsParent          bool      `db:"is_parent" json:"is_parent"`
	PrimaryLocationID string    `db:"primary_location_id" json:"primary_location_id"`
	AccessToken       string    `db:"access_token" json:"access_token"`
	RoleId            int       `db:"role_id" json:"role_id"`
	Role              string    `json:"role"`
	IsActive          bool      `db:"is_active" json:"is_active"`
	CreatedAt         time.Time `db:"created_at" json:"created_at"`
	ModifiedAt        time.Time `db:"modified_at" json:"modified_at"`
}
type GetUserResponse struct {
	UserI
	Accounts []AccountI `json:"accounts"`
}

type RoleI struct {
	Id           int       `json:"id" db:"id"`
	Role         string    `json:"role" db:"role"`
	DisplayValue string    `json:"display_value" db:"display_value"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	ModifiedAt   time.Time `db:"modified_at" json:"modified_at"`
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

type GetUserLocationsI struct {
	UserId string `db:"user_id" json:"user_id"`
}
