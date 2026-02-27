package service

import (
	"errors"
	"time"

	"go_starter_template/config"
	"go_starter_template/domain"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type authService struct {
	userRepo domain.UserRepository
	cfg      *config.Config
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func NewAuthService(userRepo domain.UserRepository, cfg *config.Config) domain.AuthService {
	return &authService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}

func (s *authService) Register(req domain.RegisterRequest) (*domain.TokenResponse, error) {
	// Check if user already exists
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Username: req.Username,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return s.generateTokens(user.ID)
}

func (s *authService) Login(req domain.LoginRequest) (*domain.TokenResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return s.generateTokens(user.ID)
}

func (s *authService) Refresh(req domain.RefreshRequest) (*domain.TokenResponse, error) {
	token, err := jwt.ParseWithClaims(req.RefreshToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(s.cfg.JWT.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, errors.New("invalid refresh token claims")
	}

	return s.generateTokens(claims.UserID)
}

func (s *authService) generateTokens(userID int64) (*domain.TokenResponse, error) {
	accessExpiry := time.Now().Add(s.cfg.JWT.AccessExpiry).Unix()
	accessClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(accessExpiry, 0)),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessStr, err := accessToken.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	// Refresh Token
	refreshExpiry := time.Now().Add(s.cfg.JWT.RefreshExpiry).Unix()
	refreshClaims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(refreshExpiry, 0)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshStr, err := refreshToken.SignedString([]byte(s.cfg.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	return &domain.TokenResponse{
		AccessToken:   accessStr,
		RefreshToken:  refreshStr,
		AccessExpiry:  accessExpiry,
		RefreshExpiry: refreshExpiry,
	}, nil
}
