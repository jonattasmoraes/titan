package domain

import (
	"context"

	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindUserById(ctx context.Context, id string) (*entities.User, error)
}
