package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
)

func startRoutes(router *gin.Engine, writer, reader *sqlx.DB) error {
	repo := repository.NewSqlxRepository(writer, reader)

	createUser := usecase.NewCreateUserUsecase(repo)
	getUserById := usecase.NewGetUserByIdUsecase(repo)
	listUsers := usecase.NewListUsersUsecase(repo)
	patchUser := usecase.NewPatchUserUsecase(repo)
	deleteUser := usecase.NewDeleteUserUsecase(repo)

	userHandlers := http.NewUserHandler(createUser, getUserById, listUsers, patchUser, deleteUser)

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
