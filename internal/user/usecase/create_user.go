package usecase

import (
	"errors"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

var ErrEmailAlreadyExists = errors.New("user with this email already exists, please try a different email")

type CreateUserUsecase struct {
	repo domain.UserRepository
}

func NewCreateUserUsecase(repo domain.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{repo: repo}
}

func (u *CreateUserUsecase) Execute(user *dto.UserDTO) (*dto.UserResponseDTO, error) {
	userExists, err := u.repo.FindUserByEmail(user.Email)

	if err != nil {
		return nil, err
	}

	if userExists != nil && userExists.Email == user.Email {
		return nil, ErrEmailAlreadyExists
	}

	createdUser, err := entities.NewUser(
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.Role,
		time.Now(),
		time.Now(),
	)
	if err != nil {
		return nil, err
	}

	newUser := &dto.UserResponseDTO{
		ID:        createdUser.ID,
		FirstName: createdUser.FirstName,
		LastName:  createdUser.LastName,
		Email:     createdUser.Email,
		Role:      createdUser.Role,
		CreateAt:  createdUser.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdateAt:  createdUser.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	err = u.repo.CreateUser(createdUser)
	if err != nil {
		return nil, err
	}

	return newUser, err
}
