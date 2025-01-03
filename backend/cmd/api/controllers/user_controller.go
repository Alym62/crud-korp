package controllers

import (
	"net/http"

	"github.com/Alym62/crud-korp/internal/dto/user"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/utils"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userUseCase usecases.UserUseCase
}

func NewUserController(useCase usecases.UserUseCase) userController {
	return userController{
		userUseCase: useCase,
	}
}

func (u *userController) Create(ctx *gin.Context) {
	var dto user.CreateUserDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	result, err := u.userUseCase.Create(dto.Email, dto.Password, dto.Position, dto.Role)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}

func (u *userController) GetById(ctx *gin.Context) {
	id, err := utils.FetchIdParamAndConvert(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := u.userUseCase.GetById(id)
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

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (u *userController) DeleteById(ctx *gin.Context) {
	id, err := utils.FetchIdParamAndConvert(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := u.userUseCase.DeleteById(id)
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

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}

func (u *userController) Update(ctx *gin.Context) {
	var dto user.UpdateUserDto

	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	id, err := utils.FetchIdParamAndConvert(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	user, err := u.userUseCase.Update(id, dto.Email, dto.Password, dto.Position, dto.Role)
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

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    user,
	})
}
