package usecase

import (
	"errors"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

var ErrEmailAlreadyExists = errors.New("a user with this email already exists, please try a different email")

type CreateUserUsecase struct {
	repo domain.UserRepository
}

func NewCreateUserUsecase(repo domain.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{repo: repo}
}

func (u *CreateUserUsecase) Execute(user *dto.UserDTO) error {
	userExists, _ := u.repo.FindUserByEmail(user.Email)
	if userExists != nil {
		return ErrEmailAlreadyExists
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
		return err
	}

	return u.repo.CreateUser(createdUser)
}
