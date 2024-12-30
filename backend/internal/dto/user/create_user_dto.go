package user

import "github.com/Alym62/crud-korp/internal/models"

type CreateUserDto struct {
	Email    string      `json:"email" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Position string      `json:"position"`
	Role     models.Role `json:"role"`
}
