package main

import (
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jonattasmoraes/titan/internal/config"
	"github.com/jonattasmoraes/titan/internal/server"
	_ "github.com/lib/pq"
)

func main() {
	writer, err := config.GetWriterSqlx()
	if err != nil {
		panic(err)
	}

	reader, err := config.GetReaderSqlx()
	if err != nil {
		panic(err)
	}

	defer writer.Close()
	defer reader.Close()

	config.StartMigrations(writer)

	server.StartServer(writer, reader)
}
