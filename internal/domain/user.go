package domain

import (
	"time"
)

type User struct {
	ID              int64                  `json:"id" gorm:"primaryKey"`
	UID             string                 `gorm:"type:uuid;default:gen_random_uuid();unique;not null"`
	Email           string                 `json:"email" gorm:"unique;index;not null"`
	PhoneNumber     string                 `json:"phone_number,omitempty"`
	Password        string                 `json:"-" gorm:"not null"`
	Username        string                 `json:"username" gorm:"unique"`
	FirstName       string                 `json:"first_name,omitempty"`
	LastName        string                 `json:"last_name,omitempty"`
	Verified        bool                   `json:"verified" gorm:"default:false"`
	Role            string                 `json:"role" gorm:"default:member"`
	Platform        string                 `json:"platform" gorm:"default:app"`
	State           string                 `json:"state" gorm:"pending"`
	Metadata        map[string]interface{} `gorm:"type:jsonb"`
	PasswordResetAt time.Time              `json:"password_reset_at"`
	CreatedAt       time.Time              `json:"created_at"`
	UpdatedAt       time.Time              `json:"updated_at"`
}

type UserRepository interface {
	GetAll() ([]User, error)
}

type UserService interface {
	GetUsers() ([]User, error)
}
