package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser_Valid(t *testing.T) {
	// Create a new user with valid parameters
	_, err := NewUser("Alice", "Johnson", "alice.johnson@example.com", "password123")
	assert.NoError(t, err)
}

func TestNewUser_MissingFirstName(t *testing.T) {
	// Attempt to create a user with missing first name
	_, err := NewUser("", "Johnson", "alice.johnson@example.com", "password123")
	assert.EqualError(t, err, ErrFirstNameIsRequired.Error())
}

func TestNewUser_InvalidEmail(t *testing.T) {
	// Attempt to create a user with an invalid email address
	_, err := NewUser("Alice", "Johnson", "invalid-email", "password123")
	assert.EqualError(t, err, ErrorValidation(err).Error())
}

func TestNewPatchUser_Valid(t *testing.T) {
	// Patch user with valid parameters
	_, err := NewPatchUser("Bob", "Smith", "bob.smith@example.com")
	assert.NoError(t, err)
}

func TestNewPatchUser_MissingAllParams(t *testing.T) {
	// Attempt to patch user with all parameters missing
	_, err := NewPatchUser("", "", "")
	assert.EqualError(t, err, ErrAtLeastOneParam.Error())
}

func TestNewPatchUser_InvalidEmail(t *testing.T) {
	// Attempt to patch user with an invalid email address
	_, err := NewPatchUser("Bob", "Smith", "invalid-email")
	assert.EqualError(t, err, ErrorValidation(err).Error())
}

func TestIsValidEmail_ValidEmail(t *testing.T) {
	// Test a valid email address
	assert.True(t, IsValidEmail("test@example.com"))
}

func TestIsValidEmail_InvalidEmail(t *testing.T) {
	// Test an invalid email address
	assert.False(t, IsValidEmail("invalid-email"))
}

func TestIsValidEmail_MissingDomain(t *testing.T) {
	// Test an email address missing domain
	assert.False(t, IsValidEmail("test@com"))
}

func TestIsValidEmail_MissingAtSymbol(t *testing.T) {
	// Test an email address missing the "@" symbol
	assert.False(t, IsValidEmail("testexample.com"))
}

func TestIsValidEmail_InvalidCharacters(t *testing.T) {
	// Test an email address with invalid characters
	assert.False(t, IsValidEmail("test@exam_ple.com"))
}

func TestIsValidEmail_ValidUKEmail(t *testing.T) {
	// Test a valid UK email address
	assert.True(t, IsValidEmail("test@example.co.uk"))
}

func TestUser_Validate_ValidUser(t *testing.T) {
	// Validate a user with valid parameters
	user := &User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alice.johnson@example.com",
		Password:  "password123",
	}
	err := user.Validate()
	assert.NoError(t, err)
}

func TestUser_Validate_MissingFirstName(t *testing.T) {
	// Validate a user missing first name
	user := &User{
		LastName: "Johnson",
		Email:    "alice.johnson@example.com",
		Password: "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrFirstNameIsRequired.Error())
}

func TestUser_Validate_MissingEmail(t *testing.T) {
	// Validate a user missing email
	user := &User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Password:  "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrEmailIsRequired.Error())
}

func TestUser_Validate_MissingPassword(t *testing.T) {
	// Validate a user missing password
	user := &User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "alice.johnson@example.com",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrPasswordIsRequired.Error())
}

func TestUser_Validate_ShortFirstName(t *testing.T) {
	// Validate a user with too short first name
	user := &User{
		FirstName: "A",
		LastName:  "Johnson",
		Email:     "alice.johnson@example.com",
		Password:  "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrFirstNameTooShort.Error())
}

func TestUser_Validate_MissingLastName(t *testing.T) {
	// Validate a user missing last name
	user := &User{
		FirstName: "Alice",
		Email:     "alice.johnson@example.com",
		Password:  "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrLastNameIsRequired.Error())
}

func TestUser_Validate_ShortLastName(t *testing.T) {
	// Validate a user with too short last name
	user := &User{
		FirstName: "Alice",
		LastName:  "J",
		Email:     "alice.johnson@example.com",
		Password:  "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrLastNameTooShort.Error())
}

func TestUser_Validate_InvalidEmail(t *testing.T) {
	// Validate a user with an invalid email address
	user := &User{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "invalid-email",
		Password:  "password123",
	}
	err := user.Validate()
	assert.EqualError(t, err, ErrorValidation(err).Error())
}

func TestUser_Patch_ValidPatch(t *testing.T) {
	// Patch a user with valid parameters
	user := &User{
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "bob.smith@example.com",
	}
	err := user.Patch()
	assert.NoError(t, err)
}

func TestUser_Patch_MissingAllParams(t *testing.T) {
	// Attempt to patch a user with all parameters missing
	user := &User{}
	err := user.Patch()
	assert.EqualError(t, err, ErrAtLeastOneParam.Error())
}

func TestUser_Patch_InvalidEmailPatch(t *testing.T) {
	// Attempt to patch a user with an invalid email address
	user := &User{
		FirstName: "Bob",
		LastName:  "Smith",
		Email:     "invalid-email",
	}
	err := user.Patch()
	assert.EqualError(t, err, ErrorValidation(err).Error())
}
