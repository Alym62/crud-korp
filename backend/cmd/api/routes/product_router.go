package routes

import (
	"database/sql"

	"github.com/Alym62/crud-korp/cmd/api/controllers"
	"github.com/Alym62/crud-korp/internal/repositories"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/gin-gonic/gin"
)

func ProductRouter(router *gin.Engine, dbConnection *sql.DB) {

	productRepository := repositories.NewProductRepository(dbConnection)

	productUseCase := usecases.NewProductUseCase(productRepository)

	productController := controllers.NewProductController(productUseCase)

	v1 := router.Group("/api/v1/product")
	v1.GET("/list", productController.GetList)
	v1.GET("/:id", productController.GetById)
	v1.POST("/create", productController.Create)
	v1.PUT("/update/:id", productController.Update)
	v1.DELETE("/delete/:id", productController.DeleteById)
}