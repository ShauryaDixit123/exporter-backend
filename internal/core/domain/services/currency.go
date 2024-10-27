package services

import "time"

type Currency struct {
	IsoCode    string
	Name       string
	Symbol     string
	IsActive   bool
	CreatedAt  time.Time
	ModifiedAt time.Time
}

type Currencies struct {
	Total     int64
	Countries []Country
}
