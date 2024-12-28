package dto

import "github.com/Alym62/crud-korp/internal/models"

type UpdateUserDto struct {
	Username string      `json:"username" binding:"required"`
	Password string      `json:"password" binding:"required"`
	Position string      `json:"position"`
	Role     models.Role `json:"role"`
}
