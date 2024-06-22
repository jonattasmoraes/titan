package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func StartServer(writer, reader *sqlx.DB) {
	router := gin.Default()

	startRoutes(router, writer, reader)

	router.Run(":8080")
}
