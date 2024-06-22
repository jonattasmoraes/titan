package infra

import (
	"github.com/jmoiron/sqlx"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type UserRepositoryPg struct {
	Db *sqlx.DB
}

func NewUserRepositoryPg(db *sqlx.DB) (domain.UserRepository, error) {
	return &UserRepositoryPg{Db: db}, nil
}

func (r *UserRepositoryPg) CreateUser(user *entities.User) error {
	query := `
	INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.Db.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPg) FindUserById(id string) (*entities.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, role, created_at, updated_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL
	`

	var user entities.User
	err := r.Db.Get(&user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
