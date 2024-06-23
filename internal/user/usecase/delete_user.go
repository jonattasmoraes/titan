package usecase

import (
	"github.com/jonattasmoraes/titan/internal/user/domain"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
)

type DeleteUserUsecase struct {
	repo domain.UserRepository
}

func NewDeleteUserUsecase(repo domain.UserRepository) *DeleteUserUsecase {
	return &DeleteUserUsecase{repo: repo}
}

func (u *DeleteUserUsecase) Execute(id string) (*dto.UserResponseDTO, error) {
	user, err := u.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}

	response := &dto.UserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		CreateAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	err = u.repo.DeleteUser(user.ID)
	if err != nil {
		return nil, err
	}

	return response, err
}
