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
	GetUserById *usecase.GetUserByIdUsecase,
	ListUsers *usecase.ListUsersUsecase,
	PatchUser *usecase.PatchUserUsecase,
	DeleteUser *usecase.DeleteUserUsecase,
) *UserHandler {
	return &UserHandler{
		createUser:  createUser,
		getUserById: GetUserById,
		listUsers:   ListUsers,
		patchUser:   PatchUser,
		deleteUser:  DeleteUser,
	}
}

// @BasePath /api/users
// @Summary Create User
// @Description POST endpoint of the application that creates new users in the database.
// @Tags Users
// @Accept json
// @Produce jsons
// @Router /users [post]
func (h *UserHandler) CreateUser(ctx *gin.Context) {
	var request dto.UserDTO

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

func (h *UserHandler) PatchUser(ctx *gin.Context) {
	var request dto.UserDTO

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
