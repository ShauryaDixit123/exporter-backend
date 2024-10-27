package services

import "time"

type Country struct {
	IsoCode     string
	Name        string
	Iso3        *string
	CallingCode uint
	IsActive    bool
	CreatedAt   time.Time
	ModifiedAt  time.Time
}

type Countries struct {
	Total     int64
	Countries []Country
}
