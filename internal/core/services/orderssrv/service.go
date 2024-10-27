package orderssrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
)

type Service struct {
	logger            logging.Logger
	purchaseOrderRepo ports.RdbmsPurchaseOrderRepoistory
	lineitemsRepo     ports.RdbmsPurchaseOrderLineItemsRepoistory
	salesOrderRepo    ports.RdbmsSalesOrderRepoistory
	workflowRepo      ports.RdbmsWorkflowRepository
}

func New(
	logger logging.Logger,
	purchaseOrderRepo ports.RdbmsPurchaseOrderRepoistory,
	lineitemsRepo ports.RdbmsPurchaseOrderLineItemsRepoistory,
	salesOrderRepo ports.RdbmsSalesOrderRepoistory,
	workflowRepo ports.RdbmsWorkflowRepository,
) *Service {
	return &Service{
		logger:            logger,
		purchaseOrderRepo: purchaseOrderRepo,
		lineitemsRepo:     lineitemsRepo,
		salesOrderRepo:    salesOrderRepo,
		workflowRepo:      workflowRepo,
	}
}

func (s *Service) CreatePurchaseOrder(
	d rdbms.PurchaseOrder,
) error {

	return nil
}
