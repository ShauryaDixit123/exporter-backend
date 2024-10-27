package rdbms

import (
	"time"
)

type Currency struct {
	IsoCode    string    `db:"iso_code"`
	Name       string    `db:"name"`
	Symbol     string    `db:"symbol"`
	IsActive   bool      `db:"is_active"`
	CreatedAt  time.Time `db:"created_at"`
	ModifiedAt time.Time `db:"modified_at"`
}
