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

//	type LoginResult struct {
//		Token string
//		User  *models.User
//	}
type LoginResult struct {
	AccessToken  string
	RefreshToken string
	User         *models.User
}

type AuthService struct {
	userRepo         *repositories.UserRepository
	refreshTokenRepo *repositories.RefreshTokenRepository
	emailService     *EmailService
}

func NewAuthService(
	userRepo *repositories.UserRepository,
	refreshTokenRepo *repositories.RefreshTokenRepository,
	emailService *EmailService,
) *AuthService {
	return &AuthService{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		emailService:     emailService,
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
	// token, err := utils.GenerateToken(
	// 	user.ID,
	// 	user.Role,
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// return &LoginResult{
	// 	Token: token,
	// 	User:  user,
	// }, nil
	accessToken, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, err
	}

	if err := s.refreshTokenRepo.DeleteByUserID(ctx, user.ID); err != nil {
		return nil, err
	}
	
	rawRefresh, refreshHash, err := utils.GenerateRefreshToken()
	if err != nil {
		return nil, err
	}
	refreshToken := &models.RefreshToken{
		UserID:    user.ID,
		TokenHash: refreshHash,
		ExpiresAt: time.Now().AddDate(0, 0, 30), // 30 วัน
	}
	if err := s.refreshTokenRepo.Create(ctx, refreshToken); err != nil {
		return nil, err
	}
	return &LoginResult{
		AccessToken:  accessToken,
		RefreshToken: rawRefresh,
		User:         user,
	}, nil
}

func (s *AuthService) RefreshToken(
	ctx context.Context,
	rawToken string,
) (string, error) {
	rawToken = strings.TrimSpace(rawToken)
	if rawToken == "" {
		return "", ErrInvalidRefreshToken
	}
	tokenHash := utils.HashRefreshToken(rawToken)
	stored, err := s.refreshTokenRepo.FindByTokenHash(ctx, tokenHash)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return "", ErrInvalidRefreshToken
	}
	if err != nil {
		return "", err
	}
	user, err := s.userRepo.FindByID(ctx, stored.UserID)
	if err != nil {
		return "", err
	}
	newAccessToken, err := utils.GenerateToken(user.ID, user.Role)
	if err != nil {
		return "", err
	}
	return newAccessToken, nil
}

func (s *AuthService) Logout(
	ctx context.Context,
	rawToken string,
) error {
	rawToken = strings.TrimSpace(rawToken)
	if rawToken == "" {
		return nil
	}
	tokenHash := utils.HashRefreshToken(rawToken)
	return s.refreshTokenRepo.DeleteByTokenHash(ctx, tokenHash)
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
