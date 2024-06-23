package domain

import (
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserById(id string) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	ListUsers(page int) ([]*entities.User, error)
	PatchUser(user *entities.User) error
}
