package controllers

import (
	"net/http"

	"github.com/Alym62/crud-korp/internal/models"
	"github.com/Alym62/crud-korp/internal/usecases"
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

func (p *productController) GetProducts(ctx *gin.Context) {
	products, err := p.productUseCase.GetProducts()
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
	var product models.Product
	err := ctx.BindJSON(&product)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	result, err := p.productUseCase.Create(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    result,
	})
}
