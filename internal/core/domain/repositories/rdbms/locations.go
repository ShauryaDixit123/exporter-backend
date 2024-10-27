package rdbms

import (
	"time"

	uuid "github.com/jackc/pgtype/ext/gofrs-uuid"
)

type LocationI struct {
	ID         uuid.UUID `db:"id" json:"id"`
	Line1      string    `db:"line1" json:"line1"`
	Line2      string    `db:"line2,omitempty" json:"line2,omitempty"`
	Area       string    `db:"area,omitempty" json:"area,omitempty"`
	City       string    `db:"city" json:"city"`
	State      string    `db:"state" json:"state"`
	CountryID  string    `db:"country_id" json:"country_id"`
	Pincode    string    `db:"pincode" json:"pincode"`
	IsActive   bool      `db:"is_active" json:"is_active"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	ModifiedAt time.Time `db:"modified_at" json:"modified_at"`
}

type CreateLocationI struct {
	Line1     string `db:"line1" json:"line1"`
	Line2     string `db:"line2,omitempty" json:"line2,omitempty"`
	Area      string `db:"area,omitempty" json:"area,omitempty"`
	City      string `db:"city" json:"city"`
	State     string `db:"state" json:"state"`
	Pincode   string `db:"pincode" json:"pincode"`
	CountryID string `db:"country_id" json:"country_id"`
}
