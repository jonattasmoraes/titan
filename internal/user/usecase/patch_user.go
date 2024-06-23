package usecase

import (
	"log"

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

	log.Printf("userExists: %v", userExists)

	if userExists != nil && userExists.Email == user.Email {
		return nil, ErrEmailAlreadyExists
	}

	updatedUser, err := entities.NewPatchUser(
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	err = u.repo.PatchUser(updatedUser)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseDTO{
		ID:        userExists.ID,
		FirstName: updatedUser.FirstName,
		LastName:  updatedUser.LastName,
		Email:     updatedUser.Email,
		Role:      userExists.Role,
		CreateAt:  userExists.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  updatedUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}
