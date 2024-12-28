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
	var dto dto.LoginDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := ac.authUseCase.Login(dto.Username)
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
			"error":   "User is not found",
		})
		return
	}

	if !user.CheckPassword(dto.Password) {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Password is not compatible",
		})
		return
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Failed to generate token",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    token,
	})
}
