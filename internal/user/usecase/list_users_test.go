package usecase_test

import (
	"testing"
	"time"

	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/oklog/ulid/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestListUsers tests the ListUsers use case.
// It verifies if the use case returns the correct users based on the provided page number.
func TestListUsers(t *testing.T) {
	// Create a new mock repository for the user repository.
	mockRepo := new(repository.MockUserRepository)

	// Create a new ListUsersUsecase with the mock repository.
	listUsersUsecase := usecase.NewListUsersUsecase(mockRepo)

	// Create a new user with a random ID.
	userId := ulid.Make().String()
	userMock := []*entities.User{
		{
			ID:        userId,
			FirstName: "John",
			LastName:  "Lennon",
			Email:     "john.lennon@example.com",
			Password:  "password",
			Role:      "user",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	// Mock the ListUsers method of the mock repository to return the user.
	mockRepo.On("ListUsers", mock.AnythingOfType("int")).Return(userMock, nil)

	// Execute the ListUsersUsecase with the page number 1.
	output, err := listUsersUsecase.Execute(1)

	// Assert that there is no error.
	assert.NoError(t, err)

	// Assert that the output is not nil and has a length of 1.
	assert.NotNil(t, output)
	assert.Len(t, output, 1)

	// Assert that the returned user is equal to the expected user.
	assert.Equal(t, userMock[0].ID, output[0].ID)
	assert.Equal(t, userMock[0].FirstName, output[0].FirstName)
	assert.Equal(t, userMock[0].LastName, output[0].LastName)
	assert.Equal(t, userMock[0].Email, output[0].Email)
	assert.Equal(t, userMock[0].Role, output[0].Role)

	// Assert that all expectations set on the mock repository have been met.
	mockRepo.AssertExpectations(t)
}
