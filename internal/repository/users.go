package repository

import (
	"context"

	"github.com/advor2102/socialnetwork/internal/models"
)

func (r *Repository) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	if err = r.db.SelectContext(ctx, &users, `
		SELECT id, user_name, email, age
		FROM users
		ORDER BY id`); err != nil {
		return nil, r.translateError(err)
	}

	return users, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	if err = r.db.GetContext(ctx, &user, `
		SELECT id, user_name, email, age
		FROM users
		WHERE id = $1`, id); err != nil {
		return models.User{}, r.translateError(err)
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (err error) {
	_, err = r.db.ExecContext(ctx, `INSERT INTO users(user_name, email, age) 
						VALUES ($1, $2, $3)`,
		user.UserName,
		user.Email,
		user.Age)
	if err != nil {
		return r.translateError(err)
	}
	return nil
}

func (r *Repository) UpdateUserByID(ctx context.Context, user models.User) (err error) {
	_, err = r.db.ExecContext(ctx, `UPDATE users SET user_name = $1, email = $2, age = $3 WHERE id = $4`,
		user.UserName,
		user.Email,
		user.Age,
		user.ID)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id int) (err error) {
	_, err = r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return r.translateError(err)
	}

	return nil
}
