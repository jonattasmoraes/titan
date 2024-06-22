package main

import (
	"os"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	"github.com/jonattasmoraes/titan/internal/config"
	"github.com/jonattasmoraes/titan/internal/server"
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

	server.StartServer(writer, reader)
}
