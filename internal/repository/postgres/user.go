package postgres

import (
	"go_starter_template/internal/domain"
)

type userRepository struct {
	db *DB
}

func NewUserRepository(db *DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetAll() ([]domain.User, error) {
	var users []domain.User
	err := r.db.Client.Find(&users).Error
	return users, err
}
