package config

import (
	"github.com/jmoiron/sqlx"
)

func GetReaderSqlx(dsn string) (*sqlx.DB, error) {
	reader, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return reader, nil
}

func GetWriterSqlx(dsn string) (*sqlx.DB, error) {
	writer, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err)
	}

	return writer, err
}
