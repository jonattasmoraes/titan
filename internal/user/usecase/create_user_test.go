package usecase_test

import (
	"testing"
	"time"

	dto "github.com/jonattasmoraes/titan/internal/user/domain/DTO"
	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestCreateUser tests the CreateUser usecase.
// It verifies if the usecase correctly creates a new user.
func TestCreateUser(t *testing.T) {
	// Create a new mock repository for the user repository.
	mockRepo := new(repository.MockUserRepository)

	// Create a new CreateUserUsecase with the mock repository.
	createUserUsecase := usecase.NewCreateUserUsecase(mockRepo)

	// Mock the FindUserByEmail method of the mock repository to return a predefined user.
	mockRepo.On("FindUserByEmail", mock.AnythingOfType("string")).Return(
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

	// Mock the CreateUser method of the mock repository to return no error.
	mockRepo.On("CreateUser", mock.AnythingOfType("*entities.User")).Return(nil)

	// Define the user DTO to be passed to the usecase.
	response := &dto.UserDTO{
		FirstName: "Peter",
		LastName:  "Parker",
		Email:     "peter.parker@example.com",
		Password:  "12345678",
		Role:      "user",
		CreateAt:  time.Now().String(),
		UpdateAt:  time.Now().String(),
	}

	// Execute the usecase with the defined user DTO.
	userDTO, err := createUserUsecase.Execute(response)

	// Assert that no error occurred during the execution.
	assert.Nil(t, err)

	// Assert that the returned user DTO matches the expected values.
	assert.Equal(t, userDTO.FirstName, response.FirstName)
	assert.Equal(t, userDTO.LastName, response.LastName)
	assert.Equal(t, userDTO.Email, response.Email)
	assert.Equal(t, userDTO.Role, response.Role)
	assert.NotEmpty(t, response.CreateAt)
	assert.NotEmpty(t, response.UpdateAt)

	// Assert that all expected methods of the mock repository were called.
	mockRepo.AssertExpectations(t)
}
