package repository

import (
	"context"
	"os"

	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/rs/zerolog"
)

func (r *Repository) CreateEmployee(ctx context.Context, employee models.Employee) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateEmployee").Logger()

	_, err = r.db.ExecContext(ctx, `INSERT INTO employees(full_name, employee_name, password, role) 
						VALUES ($1, $2, $3, $4)`,
		employee.FullName,
		employee.EmployeeName,
		employee.Password,
		employee.Role)
	if err != nil {
		logger.Err(err).Msg("error inserting employee")
		return r.translateError(err)
	}
	return nil
}

func (r *Repository) GetEmployeeByID(ctx context.Context, id int) (employee models.Employee, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetEmployeeByID").Logger()

	if err = r.db.GetContext(ctx, &employee, `
		SELECT id, full_name, employee_name, password, role, created_at, updated_at
		FROM employees
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting employee")
		return models.Employee{}, r.translateError(err)
	}

	return employee, nil
}

func (r *Repository) GetEmployeeByEmployeeName(ctx context.Context, employeeName string) (employee models.Employee, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetEmployeeByID").Logger()

	if err = r.db.GetContext(ctx, &employee, `
		SELECT id, full_name, employee_name, password, role, created_at, updated_at
		FROM employees
		WHERE employee_name = $1`, employeeName); err != nil {
		logger.Err(err).Msg("error selecting employee")
		return models.Employee{}, r.translateError(err)
	}

	return employee, nil
}
