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

// GetAllByPage recupera produtos paginados
// @Summary      Recupera produtos paginados
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        page query int false "Número da página" default(1)
// @Param        limit query int false "Limite de itens por página" default(10)
// @Success      200  {object}  []internal.models.Product
// @Failure      400  {object}  gin.H
// @Failure      500  {object}  gin.H
// @Router       /pager [get]
func (p *productController) GetAllByPage(ctx *gin.Context) {
	page, err := utils.ConverterStrToInt(ctx.DefaultQuery("page", "1"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	limit, err := utils.ConverterStrToInt(ctx.DefaultQuery("limit", "1"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	pageResponse, err := p.productUseCase.GetAllByPage(page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    pageResponse,
	})
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

func (p *productController) Update(ctx *gin.Context) {
	var dto dto.UpdateProductDto

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

	product, err := p.productUseCase.Update(id, dto.Name, dto.Description, dto.Price)
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
