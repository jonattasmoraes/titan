package usecase_test

import (
	"testing"

	"github.com/jonattasmoraes/titan/internal/user/domain/entities"
	"github.com/jonattasmoraes/titan/internal/user/infra/repository"
	"github.com/jonattasmoraes/titan/internal/user/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// TestDeleteUser tests the DeleteUser use case.
// It tests the deletion of a user by its ID.
func TestDeleteUser(t *testing.T) {
	// Create a mock UserRepository.
	mockRepo := new(repository.MockUserRepository)

	// Create a new DeleteUserUsecase with the mock repository.
	deleteUserUsecase := usecase.NewDeleteUserUsecase(mockRepo)

	// Set up the mock repository to return a User entity when FindUserById is called.
	mockRepo.On("FindUserById", mock.AnythingOfType("string")).Return(&entities.User{
		ID: "1",
	}, nil)

	// Set up the mock repository to return nil when DeleteUser is called.
	mockRepo.On("DeleteUser", mock.AnythingOfType("string")).Return(nil)

	// Execute the DeleteUser use case with the ID "1".
	_, err := deleteUserUsecase.Execute("1")

	// Assert that there is no error.
	assert.Nil(t, err)

	// Assert that all expectations set on the mock repository have been met.
	mockRepo.AssertExpectations(t)
}
