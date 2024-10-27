package workflowssrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/google/uuid"
)

type Service struct {
	logger       logging.Logger
	workflowRepo ports.RdbmsWorkflowRepository
	accountsRepo ports.RdbmsAccountsRepository
}

func New(logger logging.Logger, workflowRepo ports.RdbmsWorkflowRepository, accountsRepo ports.RdbmsAccountsRepository) *Service {
	return &Service{
		logger:       logger,
		workflowRepo: workflowRepo,
		accountsRepo: accountsRepo,
	}
}

func (s *Service) Create(m rdbms.CreateWorkflowI) error {
	id, er := s.workflowRepo.Insert(rdbms.WorkflowI{
		Name: m.Name,
		Type: m.Type,
	})
	if er != nil {
		return er
	}
	for _, v := range m.Flows {
		fl := rdbms.FlowI{
			WorkflowID:  id,
			Description: v.Description,
			Type:        v.Type,
			Title:       v.Title,
			Order:       v.Order,
			Active:      v.Active,
			TAT:         v.TAT,
		}
		flId, er := s.workflowRepo.InsertFlow(fl)
		if er != nil {
			return er
		}
		ar := make([]rdbms.FlowParamI, 0)
		for _, val := range v.FlowParams {
			ar = append(ar, rdbms.FlowParamI{
				FlowID:    flId,
				Name:      val.Name,
				Type:      val.Type,
				Mandatory: val.Mandatory,
			})
		}
		if er = s.workflowRepo.InsertFlowParams(ar); er != nil {
			return er
		}
	}
	return nil
}

func (s *Service) CreateWorkflowInstance(m rdbms.CreateWorkflowInstanceI) error {
	fls, er := s.workflowRepo.GetFlows(m.Wid)
	if er != nil {
		return er
	}
	instId := uuid.New()
	for _, v := range fls {
		ar := make([]rdbms.FlowInstanceParamI, 0)
		fid, er := s.workflowRepo.InsertFlowInstance(rdbms.FlowInstanceI{
			WorkflowID:  v.WorkflowID,
			Description: v.Description,
			Type:        v.Type,
			Title:       v.Title,
			Order:       v.Order,
			Active:      v.Active,
			TAT:         v.TAT,
			InstanceID:  instId.String(),
		})
		if er != nil {
			return er
		}
		fp, er := s.workflowRepo.GetFlowParams(v.ID)
		if er != nil {
			return er
		}
		for _, k := range fp {
			ar = append(ar, rdbms.FlowInstanceParamI{
				FlowInstanceID: fid,
				Name:           k.Name,
				Type:           k.Type,
				Mandatory:      k.Mandatory,
				Value:          nil,
			})
		}
		if er := s.workflowRepo.InsertFlowInstanceParam(ar); er != nil {
			return er
		}
	}
	if er := s.accountsRepo.InsertFlowInstanceAccount(rdbms.CreateFlowInstanceAccountI{
		AccountId:      *m.AccountId,
		FlowInstanceId: instId.String(),
	}); er != nil {
		return er
	}
	return nil
}

func (s *Service) Get(id string) ([]rdbms.GetWorkflowI, error) {
	return s.workflowRepo.Get(id)
}
func (s *Service) GetAll(of string) ([]rdbms.WorkflowI, error) {
	return s.workflowRepo.GetAll(of)
}
