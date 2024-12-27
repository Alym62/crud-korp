package main

import (
	"fmt"
	"net/http"

	"github.com/Alym62/crud-korp/cmd/api/routes"
	"github.com/Alym62/crud-korp/pkg/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// @TODO: Initializer router with gin
	router := gin.Default()
	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	// @TODO: Database connection
	dbConnection, err := db.ConnectDB()
	if err != nil {
		fmt.Println("error connection database", err)
	}

	// @TODO: All router
	routes.ProductRouter(router, dbConnection)

	router.Run(":8080")
}
