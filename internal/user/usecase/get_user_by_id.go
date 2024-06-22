package usecase

import (
	"github.com/jonattasmoraes/titan/internal/user/domain"
)

type UserResponseDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
}

type GetUserByIdUsecase struct {
	repo domain.UserRepository
}

func NewGetUserByIdUsecase(repo domain.UserRepository) *GetUserByIdUsecase {
	return &GetUserByIdUsecase{repo: repo}
}

func (u *GetUserByIdUsecase) Execute(id string) (*UserResponseDTO, error) {
	user, err := u.repo.FindUserById(id)
	if err != nil {
		return nil, err
	}

	userDTO := &UserResponseDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Role:      user.Role,
		CreateAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return userDTO, nil
}
