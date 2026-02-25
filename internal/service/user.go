package service

import (
	"go_starter_template/internal/domain"
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
