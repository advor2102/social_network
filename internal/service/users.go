package service

import (
	"errors"

	"github.com/advor2102/socialnetwork/internal/errs"
	"github.com/advor2102/socialnetwork/internal/models"
)

func (s *Service) GetAllUsers() (users []models.User, err error) {
	users, err = s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUserByID(id int) (user models.User, err error) {
	user, err = s.repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return models.User{}, errs.ErrUserNotFound
		}
		return models.User{}, err
	}

	return user, nil
}

func (s *Service) CreateUser(user models.User) (err error) {
	err = s.repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateUserByID(user models.User) (err error) {
	_, err = s.repository.GetUserByID(user.ID)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.UpdateUserByID(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteUserByID(id int) (err error) {
	_, err = s.repository.GetUserByID(id)
	if err != nil {
		if errors.Is(err, errs.ErrNotFound) {
			return errs.ErrUserNotFound
		}
		return err
	}

	err = s.repository.DeleteUserByID(id)
	if err != nil {
		return err
	}

	return nil
}
