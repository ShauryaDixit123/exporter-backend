package userssrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	logger       logging.Logger
	usersRepo    ports.RdbmsUsersRepository
	accountsRepo ports.RdbmsAccountsRepository
}

func New(logger logging.Logger,
	usersRepo ports.RdbmsUsersRepository,
	accountsRepo ports.RdbmsAccountsRepository,
) *Service {
	return &Service{
		logger:       logger,
		usersRepo:    usersRepo,
		accountsRepo: accountsRepo,
	}
}
func (s *Service) Create(
	u rdbms.CreateUserI,
) (string, error) {
	fmt.Println("dnddnnd1")
	og, er := s.usersRepo.GetById(rdbms.Id{
		Id: u.CreatedBy,
	})
	if er != nil {
		return "", er
	}
	pass, er := s.prepareUser(u)
	if er != nil {
		return "", er
	}
	u.Password = pass
	id, er := s.usersRepo.Insert(u)
	if er != nil {
		return "", er
	}
	fmt.Println("dnddnnd")
	usrac, er := s.accountsRepo.GetUserAccountById(
		u.CreatedBy,
	)
	if er != nil {
		return "", er
	}
	if u.IsParent && og.IsParent {
		newac, er := s.accountsRepo.Insert(rdbms.CreateAccountI{
			PrimaryUserID: id,
		})
		if er != nil {
			return "", er
		}
		if er := s.accountsRepo.InsertAccountUser(rdbms.CreateAccountUserI{
			AccountId: newac,
			UserId:    id,
		}); er != nil {
			return "", er
		}
		if er := s.accountsRepo.InsertAccountUser(rdbms.CreateAccountUserI{
			AccountId: usrac.Id,
			UserId:    id,
		}); er != nil {
			return "", er
		}
		return id.String(), nil
	}
	if er := s.accountsRepo.InsertAccountUser(rdbms.CreateAccountUserI{
		AccountId: usrac.Id,
		UserId:    id,
	}); er != nil {
		return "", er
	}
	return id.String(), nil
}

func (s *Service) prepareUser(u rdbms.CreateUserI) (string, error) {
	hshed, er := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if er != nil {
		return "", er
	}
	return string(hshed), nil
}

// func (s *Service) login()
//  in login we will hash user id and accountid in jwt token, so that current account be known through out application
