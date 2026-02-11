package service

import (
	"errors"
	"miniApi_BRM/internal/domain"
	"miniApi_BRM/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(req domain.CreateUserRequest) (*domain.User, error) {
	user := &domain.User{
		Name:  req.Name,
		Email: req.Email,
	}

	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetUserByID(id int) (*domain.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	return user, nil
}

func (s *UserService) UpdateUser(id int, req domain.UpdateUserRequest) (*domain.User, error) {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	if req.Name != "" {
		user.Name = req.Name
	}
	if req.Email != "" {
		user.Email = req.Email
	}

	err = s.repo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("usuario no encontrado")
	}

	return s.repo.Delete(id)
}
