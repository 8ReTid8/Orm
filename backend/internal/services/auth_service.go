package services

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"

	"backend/internal/models"
	"backend/internal/repositories"
	"backend/internal/utils"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginResult struct {
	Token string
	User  *models.User
}

type AuthService struct {
	userRepo     *repositories.UserRepository
	emailService *EmailService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	emailService *EmailService,
) *AuthService {
	return &AuthService{
		userRepo:     userRepo,
		emailService: emailService,
	}
}

func (s *AuthService) Register(

	ctx context.Context,
	email string,
	password string,

) error {
	// func (s *AuthService) Register(
	// 	ctx context.Context,
	// 	email string,
	// 	password string,
	// ) (string, error) {
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
	token, tokenHash, err := utils.GenerateVerificationToken()
	if err != nil {
		return err
	}

	expiresAt := time.Now().Add(15 * time.Minute)

	user := &models.User{
		Email:                 email,
		Password:              string(hashedPassword),
		Role:                  "user",
		IsVerified:            false,
		VerificationTokenHash: tokenHash,
		VerificationExpiresAt: &expiresAt,
	}
	if err := s.userRepo.Create(ctx, user); err != nil {
		return err
	}
	if err := s.emailService.SendVerificationEmail(
		user.Email,
		token,
	); err != nil {
		log.Printf(
			"send verification email failed: %v",
			err,
		)

		return ErrVerificationEmailFailed
	}

	return nil
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
	if !user.IsVerified {
		return nil, ErrEmailNotVerified
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

func (s *AuthService) VerifyEmail(
	ctx context.Context,
	token string,
) error {
	token = strings.TrimSpace(token)

	if token == "" {
		return ErrInvalidVerificationToken
	}

	tokenHash := utils.HashVerificationToken(token)

	user, err := s.userRepo.FindByValidVerificationTokenHash(
		ctx,
		tokenHash,
	)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrInvalidVerificationToken
	}
	if err != nil {
		return err
	}

	return s.userRepo.MarkEmailVerified(
		ctx,
		user.ID,
	)
}
