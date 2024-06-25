package repository

import (
	"errors"
	"strconv"
	"strings"
	"time"

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

// CreateUser inserts a new user into the database.
//
// Parameters:
// - user: a pointer to an entities.User struct representing the user to be created.
// Returns:
// - error: an error if the insertion operation fails, otherwise nil.
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

// FindUserById retrieves a user from the database by their ID.
//
// It takes in a single parameter, `id`, which is the ID of the user to be retrieved.
// The function returns a pointer to the User entity representing the retrieved user,
// or an error if the retrieval operation fails.
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
	found := false

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
		found = true
	}

	if !found {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// FindUserByEmail retrieves a user from the database by email.
//
// Parameters:
// - email: a string representing the email address of the user to retrieve.
// Returns:
// - *entities.User: a pointer to the User entity representing the retrieved user.
// - error: an error if the retrieval operation fails, otherwise nil.
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

// ListUsers retrieves a list of users from the database based on the provided page number.
//
// Parameters:
// - page: an integer representing the page number for pagination.
// Returns:
// - []*entities.User: a slice of User entities representing the retrieved users.
// - error: an error if the retrieval operation fails, otherwise nil.
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

// PatchUser updates the specified user in the database with the provided fields.
//
// Parameters:
// - user: a pointer to an entities.User struct representing the user to be updated.
//
// Returns:
// - error: an error if the update operation fails, otherwise nil.
func (r *repoSqlx) PatchUser(user *entities.User) error {
	var (
		query strings.Builder
		args  []interface{}
	)

	query.WriteString("UPDATE users SET ")

	argIndex := 1

	if user.FirstName != "" {
		query.WriteString("first_name = $" + strconv.Itoa(argIndex) + ", ")
		args = append(args, user.FirstName)
		argIndex++
	}

	if user.LastName != "" {
		query.WriteString("last_name = $" + strconv.Itoa(argIndex) + ", ")
		args = append(args, user.LastName)
		argIndex++
	}

	if user.Email != "" {
		query.WriteString("email = $" + strconv.Itoa(argIndex) + ", ")
		args = append(args, user.Email)
		argIndex++
	}

	updateTime := time.Now()
	query.WriteString("updated_at = $" + strconv.Itoa(argIndex))
	args = append(args, updateTime)
	argIndex++

	query.WriteString(" WHERE id = $" + strconv.Itoa(argIndex) + " AND deleted_at IS NULL")
	args = append(args, user.ID)

	_, err := r.writer.Exec(query.String(), args...)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUser apllies a date to a column teleted_at in the database.
//
// It takes in a single parameter, `id`, which is the ID of the user to be deleted.
// The function returns an error if there was a problem executing the database query.
func (r *repoSqlx) DeleteUser(id string) error {
	deletedTime := time.Now()
	query := `
	UPDATE users
	SET deleted_at = $1
	WHERE id = $2
	`

	_, err := r.writer.Exec(query, deletedTime, id)
	if err != nil {
		return err
	}

	return nil
}
