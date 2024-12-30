package routes

import (
	"database/sql"

	"github.com/Alym62/crud-korp/cmd/api/controllers"
	"github.com/Alym62/crud-korp/internal/repositories"
	"github.com/Alym62/crud-korp/internal/usecases"
	"github.com/Alym62/crud-korp/pkg/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.Engine, dbConnection *sql.DB) {

	userRepository := repositories.NewUserRepository(dbConnection)

	userUseCase := usecases.NewUserUseCase(userRepository)

	userController := controllers.NewUserController(userUseCase)

	v1 := router.Group("/api/v1/user")
	v1.Use(middlewares.CORSMiddlewares())
	v1.GET("/:id", userController.GetById)
	v1.POST("/create", userController.Create)
	v1.PUT("/update/:id", userController.Update)
	v1.DELETE("/delete/:id", userController.DeleteById)
}
