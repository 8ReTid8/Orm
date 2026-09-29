package repositories

import (
	"context"
	"time"

	"backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(
	db *gorm.DB,
) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	var user models.User

	if err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).
		Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Create(
	ctx context.Context,
	user *models.User,
) error {
	return r.db.WithContext(ctx).
		Create(user).
		Error
}

func (r *UserRepository) FindByValidVerificationTokenHash(
    ctx context.Context,
    tokenHash string,
) (*models.User, error) {
    var user models.User

    err := r.db.WithContext(ctx).
        Where("verification_token_hash = ?", tokenHash).
        Where("verification_expires_at > ?", time.Now()).
        First(&user).
        Error

    if err != nil {
        return nil, err
    }

    return &user, nil
}

func (r *UserRepository) MarkEmailVerified(
    ctx context.Context,
    userID uint,
) error {
    return r.db.WithContext(ctx).
        Model(&models.User{}).
        Where("id = ?", userID).
        Updates(map[string]any{
            "is_verified":               true,
            "verification_token_hash":   "",
            "verification_expires_at":   nil,
        }).
        Error
}