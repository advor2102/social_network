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

func (r *Repository) CreateUser(user models.User) (err error) {
	_, err = r.db.Exec(`INSERT INTO users(user_name, email, age) 
						VALUES ($1, $2, $3)`,
		user.UserName,
		user.Email,
		user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdateUserByID(user models.User) (err error) {
	_, err = r.db.Exec(`UPDATE users SET user_name = $1, email = $2, age = $3 WHERE id = $4`,
		user.UserName,
		user.Email,
		user.Age,
		user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteUserbyID(id int) (err error) {
	_, err = r.db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}

	return nil
}
