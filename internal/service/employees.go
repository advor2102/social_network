package service

import (
	"context"
	"errors"

	"github.com/advor2102/socialnetwork/internal/configs"
	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/advor2102/socialnetwork/pkg"
	"github.com/advor2102/socialnetwork/utils"
)

func (s *Service) CreateEmployee(ctx context.Context, employee models.Employee) (error error) {
	_, err := s.repository.GetEmployeeByEmployeeName(ctx, employee.EmployeeName)
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return err
		}
	} else {
		return errs.ErrEmployeeNameAlreadyExist
	}

	employee.Password, err = utils.GenerateHash(employee.Password)
	if err != nil {
		return err
	}

	if err = s.repository.CreateEmployee(ctx, employee); err != nil {
		return err
	}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, employee models.Employee)(string, error){
	empFromDB, err := s.repository.GetEmployeeByEmployeeName(ctx, employee.EmployeeName)
	if err != nil{
		if !errors.Is(err, errs.ErrNotFound) {
			return "", errs.ErrEmployeeNotFound
		}

		return "", err
	}

	employee.Password, err = utils.GenerateHash(employee.Password)
	if err != nil{
		return "", err
	}

	if empFromDB.Password != employee.Password {
		return "", errs.ErrIncorrectEmployeeNameOrPassword
	}

	token, err := pkg.GenerateToken(empFromDB.ID, configs.AppSettings.AuthParams.TtlMinutes)
	if err != nil {
		return "", err
	}

	return token, nil
}
