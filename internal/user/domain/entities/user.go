package entities

import (
	"errors"
	"fmt"
	"regexp"
	"time"

	"github.com/oklog/ulid/v2"
)

var (
	ErrAllParamsRequired   = errors.New("all params are required, please try again")
	ErrFirstNameIsRequired = errors.New("param: 'FirstName' is required, please try again")
	ErrLastNameIsRequired  = errors.New("param: 'lastName' is required, please try again")
	ErrEmailIsRequired     = errors.New("param: 'email' is required, please try again")
	ErrPasswordIsRequired  = errors.New("param: 'password' is required, please try again")
	ErrRoleIsRequired      = errors.New("param: 'role' is required, please try again")
	ErrIncorrectRole       = errors.New("param: 'role' must be 'admin', 'super' or 'user', please try again")
	ErrAtLeastOneParam     = errors.New("at least one param is required, please try again")
)

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func NewUser(id string, firstName string, lastName string, email string, password string, Role string, createdAt time.Time, updatedAt time.Time) (*User, error) {
	user := &User{
		ID:        ulid.Make().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func NewPatchUser(firstName string, lastName string, email string, password string, updatedAt time.Time) (*User, error) {
	user := &User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		UpdatedAt: time.Now(),
	}

	if err := user.Patch(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Validate() error {
	if u.FirstName == "" && u.LastName == "" && u.Email == "" && u.Password == "" && u.Role == "" {
		return ErrorValidation(ErrAllParamsRequired)
	}

	if u.FirstName == "" {
		return ErrorValidation(ErrFirstNameIsRequired)
	}

	if u.LastName == "" {
		return ErrorValidation(ErrLastNameIsRequired)
	}

	if u.Password == "" {
		return ErrorValidation(ErrPasswordIsRequired)
	}

	if u.Role == "" {
		return ErrorValidation(ErrRoleIsRequired)
	}

	if u.Role != "admin" && u.Role != "super" && u.Role != "user" {
		return ErrorValidation(ErrIncorrectRole)
	}

	if u.Email == "" {
		return ErrorValidation(ErrEmailIsRequired)
	}

	if !IsValidEmail(u.Email) {
		ErrInvalidEmail := fmt.Errorf("invalid email: %s, please try again with a valid email", u.Email)
		return ErrorValidation(ErrInvalidEmail)
	}

	return nil
}

func (r *User) Patch() error {
	if r.FirstName == "" && r.LastName == "" && r.Email == "" && r.Password == "" && r.Role == "" {
		return ErrorValidation(ErrAtLeastOneParam)
	}

	if r.Role != "" && r.Role != "admin" && r.Role != "super" && r.Role != "user" {
		return ErrorValidation(ErrIncorrectRole)
	}

	if r.Email != "" && !IsValidEmail(r.Email) {
		ErrInvalidEmail := fmt.Errorf("invalid email: %s, please try again with a valid email", r.Email)
		return ErrorValidation(ErrInvalidEmail)
	}
	return nil
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func ErrorValidation(err error) error {
	return err
}
