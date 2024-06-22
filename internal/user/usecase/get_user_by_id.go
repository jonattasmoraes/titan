package usecase

import (
	"context"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type GetUserByIdUsecase struct {
	repo domain.UserRepository
}

func NewGetUserByIdUsecase(repo domain.UserRepository) *GetUserByIdUsecase {
	return &GetUserByIdUsecase{repo: repo}
}

func (u *GetUserByIdUsecase) Execute(ctx context.Context, id string) (*entities.User, error) {
	user, err := u.repo.FindUserById(ctx, id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
