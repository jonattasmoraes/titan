package usecase

import (
	"context"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type CreateUserUsecase struct {
	repo domain.UserRepository
}

func NewCreateUserUsecase(repo domain.UserRepository) *CreateUserUsecase {
	return &CreateUserUsecase{repo: repo}
}

func (u *CreateUserUsecase) Execute(ctx context.Context, user *entities.User) error {
	user, err := entities.NewUser(user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Role)

	if err != nil {
		return err
	}

	return u.repo.CreateUser(ctx, user)
}
