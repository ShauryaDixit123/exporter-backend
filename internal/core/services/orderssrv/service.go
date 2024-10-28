package orderssrv

import (
	"exporterbackend/internal/common/constants"
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"
)

type Service struct {
	logger            logging.Logger
	purchaseOrderRepo ports.RdbmsPurchaseOrderRepoistory
	lineitemsRepo     ports.RdbmsPurchaseOrderLineItemsRepoistory
	salesOrderRepo    ports.RdbmsSalesOrderRepoistory
	// workflowRepo      ports.RdbmsWorkflowRepository
	accountsRepo    ports.RdbmsAccountsRepository
	workflowService ports.WorkflowService
}

func New(
	logger logging.Logger,
	purchaseOrderRepo ports.RdbmsPurchaseOrderRepoistory,
	lineitemsRepo ports.RdbmsPurchaseOrderLineItemsRepoistory,
	salesOrderRepo ports.RdbmsSalesOrderRepoistory,
	accountsRepo ports.RdbmsAccountsRepository,
	workflowService ports.WorkflowService,
) *Service {
	return &Service{
		logger:            logger,
		purchaseOrderRepo: purchaseOrderRepo,
		lineitemsRepo:     lineitemsRepo,
		salesOrderRepo:    salesOrderRepo,
		accountsRepo:      accountsRepo,
		workflowService:   workflowService,
	}
}

func (s *Service) CreatePurchaseOrder(
	d rdbms.CreatePurchaseOrder,
) error {
	if d.WorkflowId == nil && d.InstanceId == nil {
		acc, er := s.accountsRepo.GetById(d.AccountId)
		if er != nil {
			return er
		}
		d.WorkflowId = &acc.DefaultWorkflow
	}
	wfInstance, er := s.workflowService.AttachToWorkflow(
		rdbms.AttachWorkflowReqI{
			AccountId:        d.AccountId,
			WorkflowID:       d.WorkflowId,
			InstanceId:       d.InstanceId,
			FlowInstanceType: constants.WORKFLOW_STANDARD_INBUILT_PURCHASE_ORDER_TAG,
		},
	)
	if er != nil {
		return er
	}
	d.FlowInstanceId = wfInstance.FlowInstanceId
	d.FlowInstanceParamsId = wfInstance.FlowInstanceParamsId
	return s.createPurchaseOrder(d)
}

func (s *Service) createPurchaseOrder(
	d rdbms.CreatePurchaseOrder,
) error {
	pay := rdbms.PurchaseOrder{
		AccountId:            d.AccountId,
		PONumber:             d.PONumber,
		ShipmentTerms:        d.ShipmentTerms,
		TermsAndConditions:   d.TermsAndConditions,
		DueDate:              d.DueDate,
		PickupUserLocationId: d.PickupUserLocationId,
		DropUserLocationId:   d.DropUserLocationId,
		Status:               d.Status,
		CreatedBy:            d.CreatedBy,
		ModifiedBy:           d.ModifiedBy,
		FlowInstanceId:       d.FlowInstanceId,
		FlowInstanceParamsId: d.FlowInstanceParamsId,
		SupplierId:           d.SupplierId,
		BuyerId:              d.BuyerId, // verification required if its a buyer with valid account
	}
	id, er := s.purchaseOrderRepo.InsertOne(pay)
	if er != nil {
		return er
	}
	for i := range d.LineItems {
		d.LineItems[i].PoId = id
	}
	if er := s.lineitemsRepo.Insert(d.LineItems); er != nil {
		return er
	}
	fmt.Println(id, "uduududdu")
	return nil
}
