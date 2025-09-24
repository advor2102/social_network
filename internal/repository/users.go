package repository

import (
	"context"
	"os"

	"github.com/advor2102/socialnetwork/internal/models"
	"github.com/rs/zerolog"
)

func (r *Repository) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetAllUsers").Logger()

	if err = r.db.SelectContext(ctx, &users, `
		SELECT id, user_name, email, age
		FROM users
		ORDER BY id`); err != nil {
		logger.Err(err).Msg("error selecting users")
		return nil, r.translateError(err)
	}

	return users, nil
}

func (r *Repository) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.GetUserByID").Logger()

	if err = r.db.GetContext(ctx, &user, `
		SELECT id, user_name, email, age
		FROM users
		WHERE id = $1`, id); err != nil {
		logger.Err(err).Msg("error selecting user")
		return models.User{}, r.translateError(err)
	}

	return user, nil
}

func (r *Repository) CreateUser(ctx context.Context, user models.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.CreateUser").Logger()

	_, err = r.db.ExecContext(ctx, `INSERT INTO users(user_name, email, age) 
						VALUES ($1, $2, $3)`,
		user.UserName,
		user.Email,
		user.Age)
	if err != nil {
		logger.Err(err).Msg("error inserting user")
		return r.translateError(err)
	}
	return nil
}

func (r *Repository) UpdateUserByID(ctx context.Context, user models.User) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.UpdateUserByID").Logger()

	_, err = r.db.ExecContext(ctx, `UPDATE users SET user_name = $1, email = $2, age = $3 WHERE id = $4`,
		user.UserName,
		user.Email,
		user.Age,
		user.ID)
	if err != nil {
		logger.Err(err).Msg("error updating products")
		return r.translateError(err)
	}

	return nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id int) (err error) {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("func_name", "repository.UpdateUserByID").Logger()

	_, err = r.db.ExecContext(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		logger.Err(err).Msg("error deleting products")
		return r.translateError(err)
	}

	return nil
}
