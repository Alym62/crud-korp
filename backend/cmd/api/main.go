package main

import (
	"fmt"
	"net/http"

	"github.com/Alym62/crud-korp/cmd/api/controllers"
	"github.com/Alym62/crud-korp/internal/repository"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	dbConnection, err := db.ConnectDB()
	if err != nil {
		fmt.Println("error connection database", err)
	}

	productRepository := repository.NewProductRepository(dbConnection)

	productUseCase := usecases.NewProductUseCase(productRepository)

	productController := controllers.NewProductController(productUseCase)

	v1 := router.Group("/api/v1/product")
	v1.GET("/list", productController.GetList)
	v1.GET("/:id", productController.GetById)
	v1.POST("/create", productController.Create)
	v1.PUT("/update")
	v1.DELETE("/delete/:id", productController.DeleteById)

	router.Run(":8080")
}
