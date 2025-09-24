package models

import "time"

type Employee struct {
	ID        int    `json:"id" db:"id"`
	FullName  string `json:"full_name" db:"full_name"`
	EmployeeName  string `json:"employee_name" db:"employee_name"`
	Password  string `json:"password" db:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}