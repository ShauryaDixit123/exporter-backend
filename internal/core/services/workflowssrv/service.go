package workflowssrv

import (
	"errors"
	"exporterbackend/internal/common/constants"
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

func (s *Service) Create(m rdbms.CreateWorkflowI) (string, error) {
	id, er := s.workflowRepo.Insert(rdbms.WorkflowI{
		Name:      m.Name,
		Type:      m.Type,
		AccountId: m.AccountId,
	})
	if er != nil {
		return "", er

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
			return "", er
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
			return "", er

		}
	}
	return id.String(), nil
}

func (s *Service) CreateWorkflowInstance(m rdbms.CreateWorkflowInstanceI) (string, error) {
	wid, er := uuid.Parse(m.Wid)
	if er != nil {
		return "", er
	}
	fls, er := s.workflowRepo.GetFlows(wid)
	if er != nil {
		return "", er
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
			return "", er

		}
		fp, er := s.workflowRepo.GetFlowParams(v.ID)
		if er != nil {
			return "", er

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
			return "", er

		}
	}
	if er := s.accountsRepo.InsertFlowInstanceAccount(rdbms.CreateFlowInstanceAccountI{
		AccountId:      m.AccountId,
		FlowInstanceId: instId.String(),
	}); er != nil {
		return "", er

	}
	return instId.String(), er

}

func (s *Service) Get(id string) ([]rdbms.GetWorkflowI, error) {
	return s.workflowRepo.GetDetails(id)
}
func (s *Service) GetAll(of string) ([]rdbms.WorkflowI, error) {
	return s.workflowRepo.GetAll(of)
}

func (s *Service) AttachToWorkflow(
	d rdbms.AttachWorkflowReqI,
) (rdbms.AttachWorkflowI, error) {
	var wfinstId string
	var er error
	if d.WorkflowID != nil {
		wfinstId, er = s.CreateWorkflowInstance(rdbms.CreateWorkflowInstanceI{
			Wid:       *d.WorkflowID,
			AccountId: d.AccountId,
		})
		if er != nil {
			return rdbms.AttachWorkflowI{}, er
		}
		if er := s.workflowRepo.CreateFlowInstanceAccount(rdbms.CreateFlowInstanceAccountI{
			AccountId:      d.AccountId,
			FlowInstanceId: wfinstId,
		}); er != nil {
			return rdbms.AttachWorkflowI{}, er
		}
	}
	if d.InstanceId != nil {
		wfinstId = *d.InstanceId
	}
	getFlowinstance := rdbms.GetFlowInstance{
		InstanceId: &wfinstId,
		Type:       &d.FlowInstanceType,
	}
	flInst, er := s.workflowRepo.GetFlowInstance(getFlowinstance)
	if er != nil {
		return rdbms.AttachWorkflowI{}, er
	}
	if flInst != nil {
		return rdbms.AttachWorkflowI{
			FlowInstanceId:       &flInst.Id,
			FlowInstanceParamsId: nil,
		}, nil
	}
	flInstParams, er := s.workflowRepo.GetFlowInstanceParams(getFlowinstance)
	if er != nil {
		return rdbms.AttachWorkflowI{}, er
	}
	if flInstParams == nil {
		return rdbms.AttachWorkflowI{}, errors.New("No_Flow Instance Or Params Found For Purchase Order")
	}
	return rdbms.AttachWorkflowI{
		FlowInstanceId:       nil,
		FlowInstanceParamsId: &flInstParams[0].Id,
	}, nil

}

func (s *Service) GetInstanceAccount(
	f rdbms.GetInstanceAccount,
) ([]rdbms.CreateFlowInstanceAccountI, error) {
	return s.workflowRepo.GetInstanceAccount(f)
}

func (s *Service) UpdateFlowInstanceParam(
	f rdbms.UpdateFlowInstanceParamsI,
) error {
	return s.workflowRepo.UpdateFlowInstanceParam(f)
}

func (s *Service) UpdateFlowInstance(
	f rdbms.UpdateFlowInstanceI,
) error {
	if f.Status != nil && *f.Status == constants.STATUS_COMPLETED {
		flowParams, er := s.workflowRepo.GetFlowInstanceParams(rdbms.GetFlowInstance{
			FlowInstanceId: &f.Id,
		})
		if er != nil {
			return er
		}
		for _, val := range flowParams {
			if val.Mandatory && val.Value == nil {
				return errors.New("Mandatory_Fields Are Missing")
			}
		}
	}
	return s.workflowRepo.UpdateFlowInstance(f)
}

func (s *Service) GetInstances(
	f rdbms.GetInstancesI,
) ([]rdbms.FlowInstanceDetails, error) {
	return s.workflowRepo.GetInstances(f)
}
func (s *Service) GetFlowForAccount(
	f rdbms.GetFlowsForAccountI,
) ([]rdbms.FlowI, error) {
	return s.workflowRepo.GetFlowsForAccount(f)
}
