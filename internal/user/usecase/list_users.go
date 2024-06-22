package usecase

import (
	"errors"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
)

var (
	ErrInvalidPageNumber = errors.New("invalid page number, enter a number greater than 0")
	ErrUsersNotFound     = errors.New("users not found")
)

type ListUsersUsecase struct {
	repo domain.UserRepository
}

func NewListUsersUsecase(repo domain.UserRepository) *ListUsersUsecase {
	return &ListUsersUsecase{repo: repo}
}

func (u *ListUsersUsecase) Execute(page int) ([]*dto.UserResponseDTO, error) {
	if page < 1 {
		return nil, ErrInvalidPageNumber
	}

	users, err := u.repo.ListUsers(page)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, ErrUsersNotFound
	}

	var usersDTO []*dto.UserResponseDTO
	for _, user := range users {
		usersDTO = append(usersDTO, &dto.UserResponseDTO{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Role:      user.Role,
			CreateAt:  user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdateAt:  user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	return usersDTO, nil
}
