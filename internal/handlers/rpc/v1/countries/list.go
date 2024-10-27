package countries

import (
	"context"
	"exporterbackend/internal/core/ports"

	pb "exporterbackend/proto/gen/app/v1"
)

func (h *Handler) List(ctx context.Context, req *pb.ListCountriesRequest) (*pb.CountriesResponse, error) {

	var sortField ports.CountriesSortField = ports.COUNTRY_ISO
	switch req.SortField {
	case pb.ListCountriesRequest_SORT_FIELD_ISO:
		sortField = ports.COUNTRY_ISO
	case pb.ListCountriesRequest_SORT_FIELD_CREATED_AT:
		sortField = ports.COUNTRY_CREATED_AT
	case pb.ListCountriesRequest_SORT_FIELD_MODIFIED_AT:
		sortField = ports.COUNTRY_MODIFIED_AT
	case pb.ListCountriesRequest_SORT_FIELD_NAME:
		sortField = ports.COUNTRY_NAME
	}

	var sortOrder ports.SortOrder = ports.SORT_ASCENDING

	switch req.SortOrder {
	case pb.ListCountriesRequest_SORT_ORDER_ASC:
		sortOrder = ports.SORT_ASCENDING
	case pb.ListCountriesRequest_SORT_ORDER_DESC:
		sortOrder = ports.SORT_DESCENDING
	}

	var filters ports.CountriesFilters

	if !req.ShowDisabled {
		isActive := true
		filters = ports.CountriesFilters{
			IsActive: &isActive,
		}
	}

	var pageNumber uint = 1
	if req.PageNumber != nil {
		pageNumber = uint(*req.PageNumber)
	}

	var itemsPerPage uint = 1
	if req.ItemsPerPage != nil {
		itemsPerPage = uint(*req.ItemsPerPage)
	}

	countries, err := h.countriesService.FetchMany(
		pageNumber,
		itemsPerPage,
		req.Search,
		ports.CountriesSort{
			Field: sortField,
			Order: sortOrder,
		},
		filters,
	)

	if err != nil {
		return &pb.CountriesResponse{
			IsSuccess: false,
			Error: &pb.ErrorResponse{
				Code:    "SYSTEM_ERROR",
				Message: "There was an error when fetching countries",
			},
		}, nil
	}

	rc := make([]*pb.Country, 0)

	for _, c := range countries.Countries {
		pbCountry := mapServiceDomainToResponse(c)
		rc = append(rc, &pbCountry)
	}

	return &pb.CountriesResponse{
		IsSuccess: true,
		Data: &pb.CountriesResponse_Data{
			Pagination: &pb.CountriesResponse_Pagingation{
				CurrentPage:  uint32(pageNumber),
				ItemsPerPage: uint32(itemsPerPage),
				TotalResults: uint32(countries.Total),
			},
			Countries: rc,
		},
	}, nil
}
