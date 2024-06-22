package http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
)

type UserHandler struct {
	createUser  *usecase.CreateUserUsecase
	getUserById *usecase.GetUserByIdUsecase
}

func NewUserHandler(createUser *usecase.CreateUserUsecase, GetUserById *usecase.GetUserByIdUsecase) *UserHandler {
	return &UserHandler{
		createUser:  createUser,
		getUserById: GetUserById,
	}
}

func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var request usecase.UserDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createUser.Execute(&request); err != nil {
		if err == entities.ErrorValidation(err) {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err == usecase.ErrEmailAlreadyExists {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, request)
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	user, err := h.getUserById.Execute(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
