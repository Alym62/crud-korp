package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Alym62/crud-korp/cmd/api/routes"
	"github.com/Alym62/crud-korp/pkg/db"
	"github.com/Alym62/crud-korp/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load("../../.env")
	if err != nil {
		fmt.Println("Não foi possível carregar o arquivo .env")
		return
	}
}

func main() {
	loadEnv()

	// @TODO: Database connection and configuration for production
	var dbConnection *sql.DB

	port, err := utils.ConverterStrToInt(os.Getenv("PORT"))
	if err != nil {
		panic("Is not possible converter int")
	}

	appProduction := os.Getenv("APP_PRODUCTION")
	if appProduction == "false" {
		conn, err := db.ConnectDB(db.ConfigConnectionDB{
			Host:     os.Getenv("HOST"),
			Port:     port,
			User:     os.Getenv("USER_DB"),
			Password: os.Getenv("PASSWORD_DB"),
			DBName:   os.Getenv("DB_NAME"),
		})

		if err != nil {
			fmt.Println("error connection database", err)
		}

		dbConnection = conn
	} else {
		dbConnection = nil
	}

	if dbConnection == nil {
		panic("No connection to a database")
	}

	// @TODO: Initializer router with gin
	router := gin.Default()
	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	// @TODO: All router
	routes.ProductRouter(router, dbConnection)

	router.Run(":8080")
}
