package repository

import (
	"github.com/advor2102/socialnetwork/internal/models"
)

func (r *Repository) GetAllUsers() (users []models.User, err error) {
	if err = r.db.Select(&users, `
		SELECT id, user_name, email, age
		FROM users
		ORDER BY id`); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) GetUserByID(id int) (user models.User, err error) {
	if err = r.db.Get(&user, `
		SELECT id, user_name, email, age
		FROM users
		WHERE id = $1`, id); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *Repository) CreateUser() {

}

func (r *Repository) UpdateUserByID() {

}

func (r *Repository) DeleteUserbyID() {

}
