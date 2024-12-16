package userssrv

import (
	"exporterbackend/internal/core/domain/repositories/rdbms"
	"exporterbackend/internal/core/ports"
	"exporterbackend/pkg/logging"

	"github.com/google/uuid"
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
	u rdbms.CreateUserRequestI,
) (string, error) {
	if u.IsInvited {
		return s.createInvited(u.CreateUserI)
	}
	return s.createWithAccount(u.CreateUserI)
}

func (s *Service) prepareUser(u rdbms.CreateUserI) (rdbms.CreateUserI, error) {
	hshed, er := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if er != nil {
		return u, er
	}
	u.Password = string(hshed)
	return u, nil
}

// func (s *Service) login()
//  in login we will hash user id and accountid in jwt token, so that current account be known through out application

func (s *Service) createWithAccount(
	u rdbms.CreateUserI,
) (string, error) {
	u, er := s.prepareUser(u)
	if er != nil {
		return "", er
	}
	uid, er := s.usersRepo.Insert(u)
	if er != nil {
		return "", er
	}
	_, er = s.createAccount(rdbms.CreateAccountI{
		PrimaryUserID: uid,
	})
	if er != nil {
		return "", er
	}
	return uid.String(), nil
}

func (s *Service) createInvited(
	u rdbms.CreateUserI,
) (string, error) {
	ac, er := s.accountsRepo.GetUserAccountById(
		u.CreatedBy,
	)
	if er != nil {
		return "", er
	}
	u, er = s.prepareUser(u)
	if er != nil {
		return "", er
	}
	uid, er := s.usersRepo.Insert(u)
	if er != nil {
		return "", er
	}
	if er := s.accountsRepo.InsertAccountUser(rdbms.CreateAccountUserI{
		AccountId: ac.Id,
		UserId:    uid,
	}); er != nil {
		return "", er
	}
	return uid.String(), nil
}

func (s *Service) createAccount(
	u rdbms.CreateAccountI,
) (int, error) {
	acid, er := s.accountsRepo.Insert(rdbms.CreateAccountI{
		PrimaryUserID: u.PrimaryUserID,
		IsActive:      true,
	})
	if er != nil {
		return 0, er
	}
	return acid, s.accountsRepo.InsertAccountUser(rdbms.CreateAccountUserI{
		AccountId: acid,
		UserId:    u.PrimaryUserID,
	})
}

func (s *Service) GetUserById(f rdbms.Id) (rdbms.GetUserResponse, error) {
	uid, er := uuid.Parse(f.Id)
	if er != nil {
		return rdbms.GetUserResponse{}, er
	}
	acc, er := s.accountsRepo.GetUserAccountsById(f.Id)
	if er != nil {
		return rdbms.GetUserResponse{}, er
	}
	ud, er := s.usersRepo.GetById(uid)
	if er != nil {
		return rdbms.GetUserResponse{}, er
	}
	return rdbms.GetUserResponse{
		UserI:    ud,
		Accounts: acc,
	}, nil
}

func (s *Service) GetUsersForAccount(
	f rdbms.GetUserForAccount,
) ([]rdbms.UserI, error) {
	return s.usersRepo.GetUsersForAccount(f)
}
