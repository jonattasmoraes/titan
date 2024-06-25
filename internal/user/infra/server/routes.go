package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
)

func startRoutes(router *gin.Engine, userHandlers *http.UserHandler) error {

	userRoutes := router.Group("/api")
	{
		userRoutes.POST("/user", userHandlers.CreateUser)
		userRoutes.GET("/user/:id", userHandlers.GetUserById)
		userRoutes.GET("/users", userHandlers.ListUsers)
		userRoutes.PATCH("/user/:id", userHandlers.PatchUser)
		userRoutes.DELETE("/user/:id", userHandlers.DeleteUser)
	}

	return nil
}
