package controllers

import (
	"net/http"

	"github.com/Alym62/crud-korp/internal/dto"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/jwt"
	"github.com/gin-gonic/gin"
)

type authController struct {
	authUseCase usecases.AuthUseCase
}

func NewAuthController(useCase usecases.AuthUseCase) authController {
	return authController{
		authUseCase: useCase,
	}
}

func (ac *authController) Login(ctx *gin.Context) {
	var d dto.LoginDto
	if err := ctx.ShouldBindJSON(&d); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := ac.authUseCase.Login(d.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if user == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Usuário não encontrado",
		})
		return
	}

	if !user.CheckPassword(d.Password) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Senha incorreta. Tente novamente",
		})
		return
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Falha ao gerar token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data": dto.ResponseAuth{
			CurrentUser: dto.CurrentUserResponse{
				ID:       user.ID,
				Email:    user.Email,
				Position: user.Position,
				Role:     user.Role,
			},
			Token: token,
		},
	})
}
