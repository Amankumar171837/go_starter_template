package service

import (
	"go_starter_template/domain"
)

type userService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) domain.UserService {
	return &userService{repo: repo}
}

func (s *userService) GetUsers() ([]domain.User, error) {
	return s.repo.GetAll()
}

func (s *userService) GetUserByEmail(email string) (*domain.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *userService) CreateUser(user *domain.User) error {
	return s.repo.Create(user)
}
