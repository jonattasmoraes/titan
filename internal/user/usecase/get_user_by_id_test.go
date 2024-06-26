package usecase_test

import (
	"testing"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
)

// TestGetUserById tests the GetUserById usecase.
// It verifies if the usecase returns the correct user based on the provided ID.
func TestGetUserById(t *testing.T) {
	// Create a new mock repository for the user repository.
	mockRepo := new(repository.MockUserRepository)
	// Create a new GetUserByIdUsecase with the mock repository.
	getUserByIdUsecase := usecase.NewGetUserByIdUsecase(mockRepo)

	// Create a new user with a random ID.
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

	// Mock the FindUserById method of the mock repository to return the user.
	mockRepo.On("FindUserById", userId).Return(user, nil)

	// Execute the GetUserByIdUsecase with the user ID.
	outPut, _ := getUserByIdUsecase.Execute(userId)

	// Verify if the returned user is equal to the expected user.
	assert.Equal(t, user.ID, outPut.ID, "ID should be equal")
	assert.Equal(t, user.FirstName, outPut.FirstName, "FirstName should be equal")
	assert.Equal(t, user.LastName, outPut.LastName, "LastName should be equal")
	assert.Equal(t, user.Email, outPut.Email, "Email should be equal")
	assert.Equal(t, user.Role, outPut.Role, "Role should be equal")
}
