package repository_test

import (
	"testing"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	_ "github.com/mattn/go-sqlite3"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestDB(t *testing.T) *sqlx.DB {
	db, err := sqlx.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("Failed to connect to SQLite in-memory database: %v", err)
	}

	// Create users table
	_, err = db.Exec(`
	CREATE TABLE users (
		id TEXT PRIMARY KEY,
		first_name TEXT,
		last_name TEXT,
		email TEXT,
		password TEXT,
		role TEXT,
		created_at TIMESTAMP,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	)
	`)
	if err != nil {
		t.Fatalf("Failed to create users table: %v", err)
	}

	return db
}

func TestCreateUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSqlxRepository(db, db)

	user := &entities.User{
		ID:        "1",
		FirstName: "John",
		LastName:  "Lennon",
		Email:     "john.lennon@example.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(user)
	assert.Nil(t, err)

	userExists, err := repo.FindUserById("1")
	assert.Nil(t, err)

	assert.Equal(t, user.ID, userExists.ID)
	assert.Equal(t, user.FirstName, userExists.FirstName)
	assert.Equal(t, user.LastName, userExists.LastName)
	assert.Equal(t, user.Email, userExists.Email)
	assert.Equal(t, user.Password, userExists.Password)
	assert.Equal(t, user.Role, userExists.Role)
	assert.NotNil(t, userExists.CreatedAt)
	assert.NotNil(t, userExists.UpdatedAt)
}

func TestFindUserById(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSqlxRepository(db, db)

	userId := ulid.Make().String()
	user := &entities.User{
		ID:        userId,
		FirstName: "John",
		LastName:  "Lennon",
		Email:     "john.lennon@example.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(user)
	assert.Nil(t, err)

	foundUser, err := repo.FindUserById(userId)
	assert.Nil(t, err)

	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.FirstName, foundUser.FirstName)
	assert.Equal(t, user.LastName, foundUser.LastName)
	assert.Equal(t, user.Email, foundUser.Email)
	assert.Equal(t, user.Password, foundUser.Password)
	assert.Equal(t, user.Role, foundUser.Role)
	assert.NotNil(t, foundUser.CreatedAt)
	assert.NotNil(t, foundUser.CreatedAt)
}

func TestFindUserByEmail(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSqlxRepository(db, db)

	userId := ulid.Make().String()
	user := &entities.User{
		ID:        userId,
		FirstName: "John",
		LastName:  "Lennon",
		Email:     "john.lennon@example.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(user)
	assert.Nil(t, err)

	foundUser, err := repo.FindUserByEmail(user.Email)
	assert.Nil(t, err)

	assert.Equal(t, user.ID, foundUser.ID)
	assert.Equal(t, user.Email, foundUser.Email)
}

func TestPatchUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSqlxRepository(db, db)

	userId := ulid.Make().String()
	initialUser := &entities.User{
		ID:        userId,
		FirstName: "John",
		LastName:  "Lennon",
		Email:     "john.lennon@example.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(initialUser)
	assert.Nil(t, err)

	updatedUser := &entities.User{
		ID:        userId,
		FirstName: "Paul",
		LastName:  "McCartney",
		Email:     "paul.mccartney@example.com",
	}

	err = repo.PatchUser(updatedUser)
	assert.Nil(t, err)

	foundUser, err := repo.FindUserById(userId)
	assert.Nil(t, err)

	assert.NotEqualValues(t, initialUser.FirstName, foundUser.FirstName)
	assert.NotEqualValues(t, initialUser.LastName, foundUser.LastName)
	assert.NotEqualValues(t, initialUser.Email, foundUser.Email)
}

func TestDeleteUser(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	repo := repository.NewSqlxRepository(db, db)

	userId := ulid.Make().String()
	user := &entities.User{
		ID:        userId,
		FirstName: "John",
		LastName:  "Lennon",
		Email:     "john.lennon@example.com",
		Password:  "password",
		Role:      "user",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := repo.CreateUser(user)
	assert.Nil(t, err)

	err = repo.DeleteUser(userId)
	assert.Nil(t, err)

	_, err = repo.FindUserById(userId)
	assert.Error(t, err)
}
