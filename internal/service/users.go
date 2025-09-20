package service

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
)

func (s *Service) GetAllUsers(ctx context.Context) (users []models.User, err error) {
	users, err = s.repository.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	return users, nil
}

var (
	defaultTTL = time.Minute * 5
)

func (s *Service) GetUserByID(ctx context.Context, id int) (user models.User, err error) {
	err = s.cache.Get(ctx, fmt.Sprintf("user_%d", id), &user)
	if err == nil {
		return user, nil
	}
	// if !errors.Is(err, redis.Nil) {
	// 	return models.User{}, err
	// }

	user, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrUserNotFound
		}
		return models.User{}, err
	}

	// return user, nil

	if err = s.cache.Set(ctx, fmt.Sprintf("user_%d", user.ID), user, defaultTTL); err != nil {
		fmt.Printf("error during cache set: %v\n", err.Error())
	}
	return user, nil

}

func (s *Service) CreateUser(ctx context.Context, user models.User) (err error) {
	err = s.repository.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUserByID(ctx context.Context, user models.User) (err error) {
	_, err = s.repository.GetUserByID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.UpdateUserByID(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserByID(ctx context.Context, id int) (err error) {
	_, err = s.repository.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.DeleteUserByID(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
