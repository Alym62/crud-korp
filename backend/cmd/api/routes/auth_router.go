package routes

import (
	"database/sql"

	"github.com/Alym62/crud-korp/cmd/api/controllers"
	"github.com/Alym62/crud-korp/internal/repositories"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.Engine, dbConnection *sql.DB) {

	authRepository := repositories.NewAuthRepository(dbConnection)

	authUseCase := usecases.NewAuthUseCase(authRepository)

	authController := controllers.NewAuthController(authUseCase)

	v1 := router.Group("/auth")
	v1.Use(middlewares.CORSMiddlewares())
	v1.POST("/login", authController.Login)
}
