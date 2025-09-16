package models

type User struct {
	ID       int    `json:"id" db:"id"`
	UserName string `json:"user_name" db:"user_name"`
	Email    string `json:"email" db:"email"`
	Age      int    `json:"age" db:"age"`
}
