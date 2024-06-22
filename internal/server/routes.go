package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
)

func startRoutes(router *gin.Engine, writer, reader *sqlx.DB) {
	repo := repository.NewSqlxRepository(writer, reader)

	createUser := usecase.NewCreateUserUsecase(repo)
	getUserById := usecase.NewGetUserByIdUsecase(repo)

	userHandlers := http.NewUserHandler(createUser, getUserById)

	userRoutes := router.Group("/api")

	{
		userRoutes.POST("/user", userHandlers.CreateUser)
		userRoutes.GET("/users/:id", userHandlers.GetUserById)
	}
}
