package service

import "github.com/advor2102/socialnetwork/internal/models"

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

func (s *Service) UpdateUserByID() {

}

func (s *Service) DeleteUserbyID() {

}
