package dto

import "github.com/Alym62/crud-korp/internal/models"

type CurrentUserResponse struct {
	Email    string      `json:"email"`
	Position string      `json:"position"`
	Role     models.Role `json:"role"`
}
