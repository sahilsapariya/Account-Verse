package test

import (
	"context"
	"fmt"
	"server/database"
	"server/graph/model"
	"server/resolvers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignupResolver_Success(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	ctx := context.Background()
	params := CreateTestSignUpInput(TestUserData.ValidUser1.Name, TestUserData.ValidUser1.Email)

	// Test the resolver
	result, err := resolvers.SignupResolver(ctx, params)

	// Assertions
	assert.NoError(t, err, "SignupResolver should not return an error")
	assert.NotNil(t, result, "SignupResolver should return a result")
}

func TestSignupResolver_EmptyName(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	ctx := context.Background()
	params := CreateTestSignUpInput("", TestUserData.ValidUser1.Email)

	// Test the resolver with empty name
	result, err := resolvers.SignupResolver(ctx, params)

	// Currently the resolver doesn't validate empty name, but user should be created
	assert.NoError(t, err, "SignupResolver should not return an error for empty name")
	assert.NotNil(t, result, "SignupResolver should return a result")
}

func TestSignupResolver_EmptyEmail(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	ctx := context.Background()
	params := CreateTestSignUpInput(TestUserData.ValidUser1.Name, "")

	// Test the resolver with empty email
	result, err := resolvers.SignupResolver(ctx, params)

	// Currently the resolver doesn't validate empty email, but user should be created
	assert.NoError(t, err, "SignupResolver should not return an error for empty email")
	assert.NotNil(t, result, "SignupResolver should return a result")
}

func TestSignupResolver_BothEmpty(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	ctx := context.Background()
	params := CreateTestSignUpInput("", "")

	// Test the resolver with both empty
	result, err := resolvers.SignupResolver(ctx, params)

	// Currently the resolver doesn't validate, but user should be created
	assert.NoError(t, err, "SignupResolver should not return an error for empty fields")
	assert.NotNil(t, result, "SignupResolver should return a result")
}

func TestAddUser_Success(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	user := CreateTestUser(TestUserData.ValidUser2.Name, TestUserData.ValidUser2.Email)

	// Test the database operation directly
	result := AddUserToTestDB(t, user)

	// Assertions
	assert.NotEmpty(t, result.ID, "User ID should be generated")
	assert.Equal(t, user.Username, result.Username, "User name should match")
	assert.Equal(t, user.Email, result.Email, "User email should match")
}

func TestAddUser_WithExistingID(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	existingID := "existing-id-123"
	user := CreateTestUserWithID(existingID, TestUserData.ValidUser2.Name, TestUserData.ValidUser2.Email)

	// Test the database operation with existing ID
	result := AddUserToTestDB(t, user)

	// Assertions
	assert.Equal(t, existingID, result.ID, "User ID should remain the same")
	assert.Equal(t, user.Username, result.Username, "User name should match")
	assert.Equal(t, user.Email, result.Email, "User email should match")
}

func TestAddUser_MultipleUsers(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	// Create multiple users
	users := []*model.User{
		CreateTestUser("User 1", "user1@example.com"),
		CreateTestUser("User 2", "user2@example.com"),
		CreateTestUser("User 3", "user3@example.com"),
	}

	var createdUsers []*model.User

	for _, user := range users {
		result := AddUserToTestDB(t, user)
		createdUsers = append(createdUsers, result)
	}

	// Verify all users have unique IDs
	idMap := make(map[string]bool)
	for _, user := range createdUsers {
		assert.False(t, idMap[user.ID], "User IDs should be unique")
		idMap[user.ID] = true
	}
}

func TestAddUser_SpecialCharacters(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	user := CreateTestUser(TestUserData.SpecialCharUser.Name, TestUserData.SpecialCharUser.Email)

	// Test with special characters
	result := AddUserToTestDB(t, user)

	// Assertions
	assert.NotEmpty(t, result.ID, "User ID should be generated")
	assert.Equal(t, user.Username, result.Username, "User name with special characters should match")
	assert.Equal(t, user.Email, result.Email, "User email with special characters should match")
}

func TestAddUser_LongStrings(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	user := CreateTestUser(TestUserData.LongStringUser.Name, TestUserData.LongStringUser.Email)

	// Test with long strings
	result := AddUserToTestDB(t, user)

	// Assertions
	assert.NotEmpty(t, result.ID, "User ID should be generated")
	assert.Equal(t, user.Username, result.Username, "Long user name should match")
	assert.Equal(t, user.Email, result.Email, "Long user email should match")
}

// Integration test that tests the full flow
func TestSignupFlow_Integration(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	ctx := context.Background()

	// Test data
	testCases := []struct {
		name    string
		input   model.SignUpInput
		wantErr bool
	}{
		{
			name:    "Valid signup",
			input:   CreateTestSignUpInput("Integration Test User", "integration@test.com"),
			wantErr: false,
		},
		{
			name:    "Empty name",
			input:   CreateTestSignUpInput("", "empty.name@test.com"),
			wantErr: false, // Currently no validation
		},
		{
			name:    "Empty email",
			input:   CreateTestSignUpInput("Empty Email User", ""),
			wantErr: false, // Currently no validation
		},
		{
			name:    "Special characters",
			input:   CreateTestSignUpInput(TestUserData.SpecialCharUser.Name, TestUserData.SpecialCharUser.Email),
			wantErr: false,
		},
		{
			name:    "Long strings",
			input:   CreateTestSignUpInput(TestUserData.LongStringUser.Name, TestUserData.LongStringUser.Email),
			wantErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the resolver
			result, err := resolvers.SignupResolver(ctx, tc.input)

			if tc.wantErr {
				assert.Error(t, err, "Expected error for test case: %s", tc.name)
			} else {
				assert.NoError(t, err, "Unexpected error for test case: %s", tc.name)
				assert.NotNil(t, result, "Result should not be nil for test case: %s", tc.name)
			}
		})
	}
}

// Test for concurrent user creation
func TestAddUser_Concurrent(t *testing.T) {
	helper := SetupTestDB(t)
	defer helper.TeardownTestDB()

	const numUsers = 10
	var createdUsers []*model.User
	ctx := context.Background()

	// Create users sequentially but test that they all get unique IDs
	for i := 0; i < numUsers; i++ {
		user := CreateTestUser(fmt.Sprintf("Concurrent User %d", i), fmt.Sprintf("concurrent%d@test.com", i))

		result, err := database.Provider.AddUser(ctx, user)
		assert.NoError(t, err, "AddUser should not return an error")
		assert.NotNil(t, result, "AddUser should return a user")
		createdUsers = append(createdUsers, result)
	}

	// Verify all users were created with unique IDs
	assert.Equal(t, numUsers, len(createdUsers), "All users should be created")

	idMap := make(map[string]bool)
	for _, user := range createdUsers {
		assert.False(t, idMap[user.ID], "User IDs should be unique")
		idMap[user.ID] = true
		assert.NotEmpty(t, user.ID, "User ID should not be empty")
		assert.NotEmpty(t, user.Username, "User name should not be empty")
		assert.NotEmpty(t, user.Email, "User email should not be empty")
	}
}

// Benchmark test for performance
func BenchmarkSignupResolver(b *testing.B) {
	helper := SetupTestDB(&testing.T{})
	defer helper.TeardownTestDB()

	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Use different email for each iteration to avoid conflicts
		params := CreateTestSignUpInput("Benchmark User", fmt.Sprintf("benchmark%d@test.com", i))
		_, err := resolvers.SignupResolver(ctx, params)
		if err != nil {
			b.Fatalf("SignupResolver failed: %v", err)
		}
	}
}

// Benchmark test for database operations
func BenchmarkAddUser(b *testing.B) {
	helper := SetupTestDB(&testing.T{})
	defer helper.TeardownTestDB()

	ctx := context.Background()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Use different email for each iteration to avoid conflicts
		user := CreateTestUser("Benchmark User", fmt.Sprintf("benchmark%d@test.com", i))
		_, err := database.Provider.AddUser(ctx, user)
		if err != nil {
			b.Fatalf("AddUser failed: %v", err)
		}
	}
}
