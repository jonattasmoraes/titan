package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"

	docs "github.com/jonattasmoraes/titan/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startRoutes(router *gin.Engine, userHandlers *http.UserHandler) error {

	BasePath := "/api"
	userRoutes := router.Group(BasePath)
	{
		userRoutes.POST("/user", userHandlers.CreateUser)
		userRoutes.GET("/user/:id", userHandlers.GetUserById)
		userRoutes.GET("/users", userHandlers.ListUsers)
		userRoutes.PATCH("/user/:id", userHandlers.PatchUser)
		userRoutes.DELETE("/user/:id", userHandlers.DeleteUser)
	}

	// Swagger Docs Setup
	docs.SwaggerInfo.BasePath = BasePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return nil
}
