package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/jonattasmoraes/titan/internal/utils"
)

type UserHandler struct {
	createUser  *usecase.CreateUserUsecase
	getUserById *usecase.GetUserByIdUsecase
	listUsers   *usecase.ListUsersUsecase
	patchUser   *usecase.PatchUserUsecase
	deleteUser  *usecase.DeleteUserUsecase
}

func NewUserHandler(
	createUser *usecase.CreateUserUsecase,
	getUserById *usecase.GetUserByIdUsecase,
	listUsers *usecase.ListUsersUsecase,
	patchUser *usecase.PatchUserUsecase,
	deleteUser *usecase.DeleteUserUsecase,
) *UserHandler {
	return &UserHandler{
		createUser:  createUser,
		getUserById: getUserById,
		listUsers:   listUsers,
		patchUser:   patchUser,
		deleteUser:  deleteUser,
	}
}

// @Tags Users
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Accept  json
// @Produce  json
// @Param user body dto.UserRequestDTO true "User"
// @Success 201 {object} dto.UserResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user [post]
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var request dto.UserRequestDTO

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.createUser == nil {
		utils.SendError(ctx, http.StatusInternalServerError, "createUser is nil")
		return
	}

	user, err := h.createUser.Execute(&request)
	if err != nil {
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

	utils.SendSuccess(ctx, "create user", user, http.StatusCreated)
}

// @Tags Users
// @Summary Get user by id
// @Description Get user by id
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user/{id} [get]
func (h *UserHandler) GetUserById(ctx *gin.Context) {
	id := ctx.Param("id")

	request, err := h.getUserById.Execute(id)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			utils.SendError(ctx, http.StatusNotFound, err.Error())
			return
		}
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "get user by id", request, http.StatusOK)
}

// @Tags Users
// @Summary List users
// @Description List users
// @Accept  json
// @Produce  json
// @Param page query int true "Page number"
// @Success 200 {array} dto.UserResponseDTO
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /users [get]
func (h *UserHandler) ListUsers(ctx *gin.Context) {
	request := ctx.DefaultQuery("page", "1")

	page, err := strconv.Atoi(request)
	if err != nil {
		utils.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	users, err := h.listUsers.Execute(page)
	if err != nil {
		if err == usecase.ErrInvalidPageNumber {
			utils.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		if err == usecase.ErrUsersNotFound {
			utils.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}

		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "list users", users, http.StatusOK)
}

// @Tags Users
// @Summary Patch user
// @Description Patch user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Param user body dto.UserDTO true "User"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 404 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user/{id} [patch]
func (h *UserHandler) PatchUser(ctx *gin.Context) {
	var request dto.PatchRequestDTO

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if h.patchUser == nil {
		utils.SendError(ctx, http.StatusInternalServerError, "patchUser is nil")
		return
	}

	user := &entities.User{
		ID:        id,
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
	}

	response, err := h.patchUser.Execute(user)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			utils.SendError(ctx, http.StatusNotFound, err.Error())
			return
		}
		if err == usecase.ErrEmailAlreadyExists {
			utils.SendError(ctx, http.StatusConflict, err.Error())
			return
		}

		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "patch user", response, http.StatusOK)
}

// @Tags Users
// @Summary Delete user
// @Description Delete user
// @Accept  json
// @Produce  json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /user/{id} [delete]
func (h *UserHandler) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	response, err := h.deleteUser.Execute(id)
	if err != nil {
		if err == usecase.ErrUserNotFound {
			utils.SendError(ctx, http.StatusNotFound, err.Error())
			return
		}
		utils.SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccess(ctx, "delete user", response, http.StatusOK)
}
