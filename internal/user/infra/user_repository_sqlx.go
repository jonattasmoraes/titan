package infra

import (
	"context"

	"github.com/jmoiron/sqlx"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type UserRepositoryPg struct {
	WritterSqlx *sqlx.DB
	ReaderSqlx  *sqlx.DB
}

func NewSqlxRepository(w *sqlx.DB, r *sqlx.DB) (domain.UserRepository, error) {
	return &UserRepositoryPg{
		WritterSqlx: w,
		ReaderSqlx:  r,
	}, nil
}

func (r *UserRepositoryPg) CreateUser(ctx context.Context, user *entities.User) error {
	query := `
	INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.WritterSqlx.ExecContext(ctx, query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepositoryPg) FindUserById(ctx context.Context, id string) (*entities.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, role, created_at, updated_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL
	`

	var user entities.User
	err := r.ReaderSqlx.GetContext(ctx, &user, query, id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
