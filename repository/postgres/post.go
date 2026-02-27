package postgres

import (
	"go_starter_template/domain"
)

type postRepository struct {
	db *DB
}

func NewPostRepository(db *DB) domain.PostRepository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(post *domain.Post) error {
	return r.db.Client.Create(post).Error
}

func (r *postRepository) GetByID(id int64) (*domain.Post, error) {
	var post domain.Post
	err := r.db.Client.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *postRepository) GetAll() ([]domain.Post, error) {
	var posts []domain.Post
	err := r.db.Client.Find(&posts).Error
	return posts, err
}

func (r *postRepository) Update(post *domain.Post) error {
	return r.db.Client.Save(post).Error
}

func (r *postRepository) Delete(id int64) error {
	return r.db.Client.Delete(&domain.Post{}, id).Error
}
