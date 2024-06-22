package config

import "github.com/jmoiron/sqlx"

func GetReaderSqlx() (*sqlx.DB, error) {
	reader, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func GetWriterSqlx() (*sqlx.DB, error) {
	writer, err := sqlx.Connect("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		panic(err)
	}

	return writer, err
}
