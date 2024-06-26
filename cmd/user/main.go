package main

import (
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/titan/internal/config"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/infra/server"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	dsn := os.Getenv("DSN")

	writer, err := config.GetWriterSqlx(dsn)
	if err != nil {
		panic(err)
	}

	reader, err := config.GetReaderSqlx(dsn)
	if err != nil {
		panic(err)
	}

	defer writer.Close()
	defer reader.Close()

	config.StartMigrations(writer)

	repo := repository.NewSqlxRepository(writer, reader)

	createUser := usecase.NewCreateUserUsecase(repo)
	getUserById := usecase.NewGetUserByIdUsecase(repo)
	listUsers := usecase.NewListUsersUsecase(repo)
	patchUser := usecase.NewPatchUserUsecase(repo)
	deleteUser := usecase.NewDeleteUserUsecase(repo)

	userHandlers := http.NewUserHandler(
		createUser,
		getUserById,
		listUsers,
		patchUser,
		deleteUser)

	server.StartServer(userHandlers)
}
