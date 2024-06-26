package usecase_test

import (
	"testing"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestPatchUser tests the PatchUser usecase.
// It verifies if the usecase patches the correct user based on the provided ID.
func TestPatchUser(t *testing.T) {
	// Create a new mock repository for the user repository.
	mockRepo := new(repository.MockUserRepository)
	// Create a new PatchUserUsecase with the mock repository.
	patchUserUsecase := usecase.NewPatchUserUsecase(mockRepo)

	// Mock the FindUserById method of the mock repository to return an existing user.
	mockRepo.On("FindUserById", mock.AnythingOfType("string")).Return(
		&entities.User{
			ID:        "1",
			FirstName: "John",
			LastName:  "Lennon",
			Email:     "john.lennon@example.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}, nil)

	// Mock the PatchUser method of the mock repository to return no error.
	mockRepo.On("PatchUser", mock.AnythingOfType("*entities.User")).Return(nil)

	// Prepare the updated user data.
	updatedUser := &entities.User{
		ID:        "1",
		FirstName: "Peter",
		LastName:  "Parker",
		Email:     "peter.parker@example.com",
	}

	// Execute the PatchUserUsecase with the updated user data.
	response, err := patchUserUsecase.Execute(updatedUser)

	// Verify if the usecase returned no error and the response is not nil.
	assert.NoError(t, err)
	assert.NotNil(t, response)
	// Verify if the response matches the updated user data.
	assert.Equal(t, updatedUser.ID, response.ID)
	assert.Equal(t, updatedUser.FirstName, response.FirstName)
	assert.Equal(t, updatedUser.LastName, response.LastName)
	assert.Equal(t, updatedUser.Email, response.Email)

	// Verify if the mock repository expectations are met.
	mockRepo.AssertExpectations(t)
}
