package service

import (
	"go_starter_template/domain"
)

type postService struct {
	repo domain.PostRepository
}

func NewPostService(repo domain.PostRepository) domain.PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(post *domain.Post) error {
	// Business logic: You could validate post title or body here
	return s.repo.Create(post)
}

func (s *postService) GetPost(id int64) (*domain.Post, error) {
	return s.repo.GetByID(id)
}

func (s *postService) GetPosts() ([]domain.Post, error) {
	return s.repo.GetAll()
}

func (s *postService) UpdatePost(post *domain.Post) error {
	return s.repo.Update(post)
}

func (s *postService) DeletePost(id int64) error {
	return s.repo.Delete(id)
}
