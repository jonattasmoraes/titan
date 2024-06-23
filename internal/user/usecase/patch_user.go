package usecase

import (
	"errors"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type PatchUserUsecase struct {
	repo domain.UserRepository
}

func NewPatchUserUsecase(repo domain.UserRepository) *PatchUserUsecase {
	return &PatchUserUsecase{repo: repo}
}

func (u *PatchUserUsecase) Execute(user *entities.User) (*dto.UserResponseDTO, error) {
	userExists, err := u.repo.FindUserById(user.ID)
	if err != nil {
		return nil, err
	}

	if userExists == nil {
		return nil, errors.New("user not found")
	}

	if userExists.Email == user.Email {
		return nil, ErrEmailAlreadyExists
	}

	updatedUser := &entities.User{
		ID:        userExists.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  userExists.Password,
		Role:      userExists.Role,
		CreatedAt: userExists.CreatedAt,
		UpdatedAt: time.Now(),
	}

	err = u.repo.PatchUser(updatedUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseDTO{
		ID:        updatedUser.ID,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Email:     updatedUser.Email,
		Role:      updatedUser.Role,
		CreateAt:  updatedUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  updatedUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
