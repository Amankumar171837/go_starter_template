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

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Client.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *domain.User) error {
	return r.db.Client.Create(user).Error
}
