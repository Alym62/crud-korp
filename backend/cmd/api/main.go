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
	env := os.Getenv("APP_PRODUCTION")
	if env != "true" {
		err := godotenv.Load("../../.env")
		if err != nil {
			fmt.Println("Não foi possível carregar o arquivo .env")
		} else {
			fmt.Println("Arquivo .env carregado com sucesso")
		}
	}
}

func main() {
	loadEnv()

	// @TODO: Database connection and configuration for production
	var dbConnection *sql.DB

	port, err := utils.ConverterStrToInt(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Is not possible converter int")
	}

	appProduction := os.Getenv("APP_PRODUCTION")
	if appProduction != "" {
		conn, err := db.ConnectDB(db.ConfigConnectionDB{
			Host:     os.Getenv("DB_HOST"),
			Port:     port,
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
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
	routes.UserRouter(router, dbConnection)
	routes.AuthRouter(router, dbConnection)

	router.Run(":8080")
}
