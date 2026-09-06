package repositories

import (
	"context"

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