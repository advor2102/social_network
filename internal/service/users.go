package service

import "github.com/advor2102/socialnetwork/internal/models"

func (s *Service) GetAllUsers() (users []models.User, err error) {
	users, err = s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) GetUserByID() {

}

func (s *Service) CreateUser() {

}

func (s *Service) UpdateUserByID() {

}

func (s *Service) DeleteUserbyID() {

}
