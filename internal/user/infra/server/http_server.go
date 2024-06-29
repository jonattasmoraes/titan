package server

import (
	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/titan/internal/user/infra/http"
)

func StartServer(userHandlers *http.UserHandler) {
	router := gin.Default()

	startRoutes(router, userHandlers)

	router.Run(":8080")
}
