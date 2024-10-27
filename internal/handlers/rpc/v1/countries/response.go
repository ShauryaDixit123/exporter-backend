package countries

import (
	"exporterbackend/internal/core/domain/services"
	"time"

	pb "exporterbackend/proto/gen/app/v1"
)

func mapServiceDomainToResponse(
	c services.Country,
) pb.Country {
	return pb.Country{
		IsoCode:     c.IsoCode,
		Name:        c.Name,
		CallingCode: uint32(c.CallingCode),
		Iso_3:       c.Iso3,
		IsActive:    c.IsActive,
		CreatedAt:   c.CreatedAt.Format(time.RFC3339),
		ModifiedAt:  c.ModifiedAt.Format(time.RFC3339),
	}
}
