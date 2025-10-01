package service

import (
	"context"
	"errors"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
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

	employee.Role = models.RoleUser

	if err = s.repository.CreateEmployee(ctx, employee); err != nil {
		return err
	}

	return nil
}

func (s *Service) Authenticate(ctx context.Context, employee models.Employee) (int, models.Role, error) {
	empFromDB, err := s.repository.GetEmployeeByEmployeeName(ctx, employee.EmployeeName)
	if err != nil {
		if !errors.Is(err, errs.ErrNotFound) {
			return 0, "", errs.ErrEmployeeNotFound
		}

		return 0, "", err
	}

	employee.Password, err = utils.GenerateHash(employee.Password)
	if err != nil {
		return 0, "", err
	}

	if empFromDB.Password != employee.Password {
		return 0, "", errs.ErrIncorrectEmployeeNameOrPassword
	}

	return empFromDB.ID, empFromDB.Role, nil
}
