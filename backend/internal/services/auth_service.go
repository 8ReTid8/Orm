package services

import (
	"context"
	"errors"
	"strings"

	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// var (
// 	ErrEmailAlreadyUsed   = errors.New("email already used")
// 	ErrInvalidCredentials = errors.New("invalid credentials")
// )

type LoginResult struct {
	Token string
	User  *models.User
}

type AuthService struct {
	userRepo *repositories.UserRepository
}

func NewAuthService(
	userRepo *repositories.UserRepository,
) *AuthService {
	return &AuthService{
		userRepo: userRepo,
	}
}

func (s *AuthService) Register(
	ctx context.Context,
	email string,
	password string,
) error {
	email = strings.ToLower(strings.TrimSpace(email))

	_, err := s.userRepo.FindByEmail(ctx, email)

	if err == nil {
		return ErrEmailAlreadyUsed
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:    email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	return s.userRepo.Create(ctx, user)
}

func (s *AuthService) Login(
	ctx context.Context,
	email string,
	password string,
) (*LoginResult, error) {
	email = strings.ToLower(strings.TrimSpace(email))

	user, err := s.userRepo.FindByEmail(ctx, email)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredentials
	}
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	); err != nil {
		return nil, ErrInvalidCredentials
	}

	token, err := utils.GenerateToken(
		user.ID,
		user.Role,
	)
	if err != nil {
		return nil, err
	}

	return &LoginResult{
		Token: token,
		User:  user,
	}, nil
}