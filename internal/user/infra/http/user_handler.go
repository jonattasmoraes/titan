package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/jonattasmoraes/titan/internal/utils"
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
	var request dto.UserDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.createUser.Execute(&request); err != nil {
		if err == entities.ErrorValidation(err) {
			utils.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err == usecase.ErrEmailAlreadyExists {
			utils.SendError(ctx, http.StatusConflict, err.Error())
			return
		}

		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "create user", nil, http.StatusCreated)
}

func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	request, err := h.getUserById.Execute(id)
	if err != nil {
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "get user by id", request, http.StatusOK)
}
