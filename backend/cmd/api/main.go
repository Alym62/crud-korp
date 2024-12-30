package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Alym62/crud-korp/cmd/api/routes"
	"github.com/Alym62/crud-korp/pkg/db"
	"github.com/Alym62/crud-korp/pkg/middlewares"
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

	// Conexão com o banco de dados
	var dbConnection *sql.DB

	port, err := utils.ConverterStrToInt(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err.Error())
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
			fmt.Println("Erro ao tentar conectar ao banco de dados.", err)
		}

		dbConnection = conn
	} else {
		dbConnection = nil
	}

	if dbConnection == nil {
		panic("Nenhuma conexão encontrada com o banco de dados.")
	}

	// Inicializando o Gin
	router := gin.Default()

	router.Use(middlewares.CORSMiddlewares())

	router.GET("/healthy", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	})

	// Todas as rotas
	routes.ProductRouter(router, dbConnection)
	routes.UserRouter(router, dbConnection)
	routes.AuthRouter(router, dbConnection)

	router.Run(":8080")
}
