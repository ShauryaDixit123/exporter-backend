package ports

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
)

type RdbmsPurchaseOrderRepoistory interface {
	InsertOne(
		data rdbms.PurchaseOrder,
	) error
	SelectOne(id string) (rdbms.PurchaseOrder, error)
	UpdateOne(
		id string,
		data rdbms.PurchaseOrder,
	) error
	UpdateStatus(
		id string,
		status string,
		rejectionReason *string,
	) error
	DeleteOne(id string) error
	CheckPOExistsByCode(code string) (bool, error)
	Count(
		search OrdersSearch,
		filters OrdersFilters,
		account_id int,
	) (int64, error)
	SelectMany(
		limit uint,
		offset uint,
		search OrdersSearch,
		sort OrdersSort,
		filters OrdersFilters,
		account_id int,
	) ([]rdbms.PurchaseOrder, error)
}

type RdbmsPurchaseOrderLineItemsRepoistory interface {
	Insert(
		data []rdbms.OrderLineItems,
	) error
	GetPOLineItems(
		id string,
		soID *string,
		og *bool,
	) ([]rdbms.OrderLineItems, error)
	DeletePOLineItems(id string, by string) error
	FetchLastInsertedLIForSOByRefId(
		id string,
	) (rdbms.OrderLineItems, error)
	GetLIBySOid(
		id string,
	) ([]rdbms.OrderLineItems, error)
	DeleteLineItemById(
		id string,
	) error
	UpdateLineItemById(
		id string,
		bq *int, // BATCH COUNT
		quan *int, // QUANTITY
		itemCode *string,
		description *string,
		deliveryDate *string,
	) error
	GetLineItemsBySOId(
		id string,
	) ([]rdbms.OrderLineItems, error)
	ExpiredLineItems() ([]rdbms.OrderLineItems, error)
}

type RdbmsSalesOrderRepoistory interface {
	Insert(
		data rdbms.SalesOrder,
	) error
	SelectByPOId(
		id string,
	) ([]rdbms.SalesOrder, error)
	CheckSOExistsByCode(code string) (bool, error)
	SelectBySOId(
		id string,
	) (rdbms.SalesOrder, error)
	DeleteById(
		id string,
	) error
	UpdateBySOId(
		id string,
		data rdbms.SalesOrder,
	) error
	Count(
		search OrdersSearch,
		filters OrdersFilters,
		account_id int,
	) (int64, error)
	SelectMany(
		limit uint,
		offset uint,
		search OrdersSearch,
		sort OrdersSort,
		filters OrdersFilters,
		account_id int,
	) ([]rdbms.SalesOrder, error)
	UpdateStatus(
		id string,
		status string,
		reason *string,
	) error
	FetchAllExpired() ([]rdbms.SalesOrder, error)
}
type OrdersService interface {
	// CreatePO(data services.PurchaseOrder) (string, error)
	// GetPODetails(id string) (services.PurchaseOrder, error)
	// UpdatePO(data services.PurchaseOrder) error
	// UpdateStatus(id string, isSO bool, status string, reason *string) error
	// InitializeSalesOrder(
	// 	poId string,
	// ) (services.CreateSalesOrder, error)
	// CreateSalesOrder(
	// 	so services.SalesOrder,
	// 	li []services.OrderLineItems,
	// ) (string, error)
	// GetSO(
	// 	soid string,
	// ) (services.CreateSalesOrder, error)
	// UpdateSalesOrder(
	// 	soid string,
	// 	so services.SalesOrder,
	// 	li []services.OrderLineItems,
	// ) error
	// OrdersCheck() error

	// FetchPOOrders(
	// 	limit uint,
	// 	offset uint,
	// 	searchTerm *string,
	// 	sort OrdersSort,
	// 	filters OrdersFilters,
	// 	account_id int,
	// 	status *string,
	// ) (services.POOrders, error)
	// FetchSOOrders(
	// 	limit uint,
	// 	offset uint,
	// 	searchTerm *string,
	// 	sort OrdersSort,
	// 	filters OrdersFilters,
	// 	account_id int,
	// 	status *string,
	// ) (services.SOOrders, error)
	// GenrateDocument(
	// 	id string,
	// 	isSO bool,
	// ) ([]byte, error)
	CreatePurchaseOrder(
		d rdbms.PurchaseOrder,
	) error

	// AddSOTrackingDetails(isContainerTracking bool, soId, containerNoOrMBLNo, liner string) error
	// GetSOTrackingDetails(soId string) ([]shipsgo.TrackingResponse, error)
	// AddBookingAgent(soId string, bookingAgent int, bookingStatus string, customerName string) error
}

type OrderSortField int

const (
	ORDER_CREATED_BY  OrderSortField = iota
	ORDER_MODIFIED_BY OrderSortField = iota
	ORDER_NAME        OrderSortField = iota
)

type OrdersSort struct {
	Field OrderSortField
	Order SortOrder
}

type OrdersSearch struct {
	Id   int
	Name *string
}

type OrdersFilters struct {
	Status *string
	User   string
}
