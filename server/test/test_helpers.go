package test

import (
	"context"
	"os"
	"server/database"
	"server/database/providers/sql"
	"server/graph/model"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

// Global mutex to ensure thread-safe database initialization
var dbMutex sync.Mutex

// TestDBHelper provides database setup and teardown for tests
type TestDBHelper struct {
	originalDBType string
	originalDBName string
}

// SetupTestDB initializes an in-memory SQLite database for testing
func SetupTestDB(t *testing.T) *TestDBHelper {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	helper := &TestDBHelper{
		originalDBType: os.Getenv("DB_TYPE"),
		originalDBName: os.Getenv("DB_NAME"),
	}

	// Set test environment variables
	os.Setenv("DB_TYPE", "sqlite")
	os.Setenv("DB_NAME", ":memory:") // Use in-memory SQLite for tests

	// Initialize test database
	provider, err := sql.NewProvider()
	require.NoError(t, err, "Failed to create test database provider")

	database.Provider = provider
	return helper
}

// TeardownTestDB cleans up the test database
func (h *TestDBHelper) TeardownTestDB() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	// Restore original environment variables
	if h.originalDBType != "" {
		os.Setenv("DB_TYPE", h.originalDBType)
	} else {
		os.Unsetenv("DB_TYPE")
	}

	if h.originalDBName != "" {
		os.Setenv("DB_NAME", h.originalDBName)
	} else {
		os.Unsetenv("DB_NAME")
	}
}

// CreateTestUser creates a test user with default values
func CreateTestUser(name, email string) *model.User {
	return &model.User{
		GivenName: &name,
		Email:     email,
	}
}

// CreateTestUserWithID creates a test user with a specific ID
func CreateTestUserWithID(id, name, email string) *model.User {
	return &model.User{
		ID:        id,
		GivenName: &name,
		Email:     email,
	}
}

// CreateTestSignUpInput creates a test signup input
func CreateTestSignUpInput(name, email string) model.SignUpInput {
	return model.SignUpInput{
		GivenName: &name,
		Email:     email,
	}
}

// AddUserToTestDB adds a user to the test database and returns the created user
func AddUserToTestDB(t *testing.T, user *model.User) *model.User {
	ctx := context.Background()
	result, err := database.Provider.AddUser(ctx, user)
	require.NoError(t, err, "Failed to add user to test database")
	require.NotNil(t, result, "AddUser should return a user")
	return result
}

// TestUserData contains common test user data
var TestUserData = struct {
	ValidUser1 struct {
		Name  string
		Email string
	}
	ValidUser2 struct {
		Name  string
		Email string
	}
	SpecialCharUser struct {
		Name  string
		Email string
	}
	LongStringUser struct {
		Name  string
		Email string
	}
}{
	ValidUser1: struct {
		Name  string
		Email string
	}{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	},
	ValidUser2: struct {
		Name  string
		Email string
	}{
		Name:  "Jane Smith",
		Email: "jane.smith@example.com",
	},
	SpecialCharUser: struct {
		Name  string
		Email string
	}{
		Name:  "José María García-López",
		Email: "jose.maria+test@example.com",
	},
	LongStringUser: struct {
		Name  string
		Email string
	}{
		Name:  "This is a very long name that might cause issues with database storage if not handled properly",
		Email: "this.is.a.very.long.email.address.that.might.cause.issues@very-long-domain-name-example.com",
	},
}
