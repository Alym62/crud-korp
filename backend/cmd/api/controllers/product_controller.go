package controllers

import (
	"net/http"

	"github.com/Alym62/crud-korp/internal/dto"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/utils"
	"github.com/gin-gonic/gin"
)

type productController struct {
	productUseCase usecases.ProductUseCase
}

func NewProductController(useCase usecases.ProductUseCase) productController {
	return productController{
		productUseCase: useCase,
	}
}

func (p *productController) GetList(ctx *gin.Context) {
	products, err := p.productUseCase.GetList()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    products,
	})
}

func (p *productController) Create(ctx *gin.Context) {
	var dto dto.CreateProductDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	result, err := p.productUseCase.Create(dto.Name, dto.Description, dto.Price)
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

func (p *productController) GetById(ctx *gin.Context) {
	id, err := utils.FetchIdParamAndConvert(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	product, err := p.productUseCase.GetById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Product is not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}

func (p *productController) DeleteById(ctx *gin.Context) {
	id, err := utils.FetchIdParamAndConvert(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	product, err := p.productUseCase.DeleteById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	if product == nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"success": false,
			"error":   "Product is not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    product,
	})
}
