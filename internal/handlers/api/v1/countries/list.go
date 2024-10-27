package countries

import (
	"exporterbackend/internal/core/ports"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) List(ctx *gin.Context) {
	var query listRequest
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var sortField ports.CountriesSortField = ports.COUNTRY_ISO

	if query.SortField != nil {
		switch *query.SortField {
		case "iso":
			sortField = ports.COUNTRY_ISO
		case "createdat":
			sortField = ports.COUNTRY_CREATED_AT
		case "modifiedat":
			sortField = ports.COUNTRY_MODIFIED_AT
		case "name":
			sortField = ports.COUNTRY_NAME
		}
	}

	var sortOrder ports.SortOrder = ports.SORT_ASCENDING

	if query.SortOrder != nil {
		switch *query.SortOrder {
		case "asc":
			sortOrder = ports.SORT_ASCENDING
		case "desc":
			sortOrder = ports.SORT_DESCENDING
		}
	}

	var filters ports.CountriesFilters

	if !query.ShowDisabled {
		isActive := true
		filters = ports.CountriesFilters{
			IsActive: &isActive,
		}
	}

	var pageNumber uint = 1
	if query.PageNumber != nil {
		pageNumber = *query.PageNumber
	}

	var itemsPerPage uint = 1
	if query.ItemsPerPage != nil {
		itemsPerPage = *query.ItemsPerPage
	}

	countries, err := h.countriesService.FetchMany(
		pageNumber,
		itemsPerPage,
		query.SearchTerm,
		ports.CountriesSort{
			Field: sortField,
			Order: sortOrder,
		},
		filters,
	)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	rc := make([]country, 0)

	for _, c := range countries.Countries {
		rc = append(rc, mapServiceDomainToResponse(c))
	}

	response := countriesResponse{
		CurrentPage:  pageNumber,
		ItemsPerPage: itemsPerPage,
		TotalResults: countries.Total,
		Countries:    rc,
	}

	ctx.JSON(http.StatusOK, response)
}

type listRequest struct {
	PageNumber   *uint   `form:"pn" binding:"omitempty,numeric"`
	ItemsPerPage *uint   `form:"ipp" binding:"omitempty,numeric"`
	SortField    *string `form:"sf" binding:"omitempty,alpha,oneof=iso createdat modifiedat name"`
	SortOrder    *string `form:"so" binding:"omitempty,alpha,oneof=asc desc"`
	SearchTerm   *string `form:"s"`
	ShowDisabled bool    `form:"sd" binding:"omitempty"`
}
