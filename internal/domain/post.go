package domain

import (
	"time"
)

type Post struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	Body      string    `json:"body" gorm:"type:text;not null"`
	UserID    int64     `json:"user_id" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepository interface {
	Create(post *Post) error
	GetByID(id int64) (*Post, error)
	GetAll() ([]Post, error)
	Update(post *Post) error
	Delete(id int64) error
}

type PostService interface {
	CreatePost(post *Post) error
	GetPost(id int64) (*Post, error)
	GetPosts() ([]Post, error)
	UpdatePost(post *Post) error
	DeletePost(id int64) error
}
