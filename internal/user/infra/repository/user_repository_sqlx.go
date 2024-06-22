package repository

import (
	"github.com/jmoiron/sqlx"

	"github.com/jonattasmoraes/titan/internal/user/domain"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
)

type repoSqlx struct {
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlxRepository(writer, reader *sqlx.DB) domain.UserRepository {
	return &repoSqlx{writer: writer, reader: reader}
}

func (r *repoSqlx) CreateUser(user *entities.User) error {
	query := `
	INSERT INTO users (id, first_name, last_name, email, password, role, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
	`

	_, err := r.writer.Exec(query, user.ID, user.FirstName, user.LastName, user.Email, user.Password, user.Role, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

func (r *repoSqlx) FindUserById(id string) (*entities.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, role, created_at, updated_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL
	`

	rows, err := r.reader.Query(query, id)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user entities.User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r *repoSqlx) FindUserByEmail(email string) (*entities.User, error) {
	query := `
	SELECT id, first_name, last_name, email, password, role, created_at, updated_at
	FROM users
	WHERE email = $1 AND deleted_at IS NULL
	`

	rows, err := r.reader.Query(query, email)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var user entities.User

	for rows.Next() {
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (r *repoSqlx) ListUsers(page int) ([]*entities.User, error) {
	offset := (page - 1) * 10

	query := `
	SELECT id, first_name, last_name, email, role, created_at, updated_at
	FROM users
	WHERE deleted_at IS NULL
	LIMIT 10 OFFSET $1
	`

	rows, err := r.writer.Query(query, offset)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []*entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}
