package rdbms

import (
	"database/sql"
	"time"
)

type Country struct {
	IsoCode     string         `db:"iso_code"`
	Name        string         `db:"name"`
	Iso3        sql.NullString `db:"iso_3"`
	CallingCode uint           `db:"calling_code"`
	IsActive    bool           `db:"is_active"`
	CreatedAt   time.Time      `db:"created_at"`
	ModifiedAt  time.Time      `db:"modified_at"`
}
