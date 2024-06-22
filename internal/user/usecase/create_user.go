package usecase

import (
	"errors"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type UserDTO struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreateAt  string `json:"create_at"`
	UpdateAt  string `json:"update_at"`
	DeletedAt string `json:"deleted_at"`
}

var ErrEmailAlreadyExists = errors.New("a user with this email already exists, please try a different email")

type CreateUserUsecase struct {
	repo domain.UserRepository
}

func NewCreateUserUsecase(repo domain.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{repo: repo}
}

func (u *CreateUserUsecase) Execute(user *UserDTO) error {
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
