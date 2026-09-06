package dto

import "backend/internal/models"

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type AuthUserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type LoginResponse struct {
	Token string           `json:"token"`
	User  AuthUserResponse `json:"user"`
}

func ToAuthUserResponse(
	user models.User,
) AuthUserResponse {
	return AuthUserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  user.Role,
	}
}