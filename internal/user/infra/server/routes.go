package server

import (
	"github.com/gin-gonic/gin"
	docs "github.com/jonattasmoraes/titan/docs"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startRoutes(router *gin.Engine, userHandlers *http.UserHandler) {
	docs.SwaggerInfo.BasePath = "/api"
	userRoutes := router.Group("/api")
	{
		userRoutes.POST("/user", userHandlers.CreateUser)
		userRoutes.GET("/user/:id", userHandlers.GetUserById)
		userRoutes.GET("/users", userHandlers.ListUsers)
		userRoutes.PATCH("/user/:id", userHandlers.PatchUser)
		userRoutes.DELETE("/user/:id", userHandlers.DeleteUser)
		userRoutes.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.NewHandler()))
	}
}
