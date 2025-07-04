package test

import (
	"os"
	"testing"
)

// TestConfig holds configuration for tests
type TestConfig struct {
	DBType       string
	DBName       string
	LogLevel     string
	TestTimeout  string
	CleanupAfter bool
}

// DefaultTestConfig returns the default test configuration
func DefaultTestConfig() *TestConfig {
	return &TestConfig{
		DBType:       "sqlite",
		DBName:       ":memory:",
		LogLevel:     "error", // Reduce log noise in tests
		TestTimeout:  "30s",
		CleanupAfter: true,
	}
}

// SetupTestEnvironment sets up the test environment with the given config
func SetupTestEnvironment(t *testing.T, config *TestConfig) func() {
	// Store original values
	originalDBType := os.Getenv("DB_TYPE")
	originalDBName := os.Getenv("DB_NAME")
	originalLogLevel := os.Getenv("LOG_LEVEL")

	// Set test values
	os.Setenv("DB_TYPE", config.DBType)
	os.Setenv("DB_NAME", config.DBName)
	os.Setenv("LOG_LEVEL", config.LogLevel)

	// Return cleanup function
	return func() {
		if config.CleanupAfter {
			// Restore original values
			if originalDBType != "" {
				os.Setenv("DB_TYPE", originalDBType)
			} else {
				os.Unsetenv("DB_TYPE")
			}

			if originalDBName != "" {
				os.Setenv("DB_NAME", originalDBName)
			} else {
				os.Unsetenv("DB_NAME")
			}

			if originalLogLevel != "" {
				os.Setenv("LOG_LEVEL", originalLogLevel)
			} else {
				os.Unsetenv("LOG_LEVEL")
			}
		}
	}
}

// TestScenarios defines common test scenarios
var TestScenarios = struct {
	ValidSignup struct {
		Name        string
		Email       string
		Description string
	}
	InvalidEmail struct {
		Name        string
		Email       string
		Description string
	}
	EmptyFields struct {
		Name        string
		Email       string
		Description string
	}
	SpecialChars struct {
		Name        string
		Email       string
		Description string
	}
	LongStrings struct {
		Name        string
		Email       string
		Description string
	}
}{
	ValidSignup: struct {
		Name        string
		Email       string
		Description string
	}{
		Name:        "Test User",
		Email:       "test@example.com",
		Description: "Valid signup with normal data",
	},
	InvalidEmail: struct {
		Name        string
		Email       string
		Description string
	}{
		Name:        "Test User",
		Email:       "invalid-email",
		Description: "Signup with invalid email format",
	},
	EmptyFields: struct {
		Name        string
		Email       string
		Description string
	}{
		Name:        "",
		Email:       "",
		Description: "Signup with empty fields",
	},
	SpecialChars: struct {
		Name        string
		Email       string
		Description string
	}{
		Name:        "José María García-López",
		Email:       "jose.maria+test@example.com",
		Description: "Signup with special characters",
	},
	LongStrings: struct {
		Name        string
		Email       string
		Description string
	}{
		Name:        "This is a very long name that might cause issues with database storage if not handled properly and exceeds normal length limits",
		Email:       "this.is.a.very.long.email.address.that.might.cause.issues.with.database.storage@very-long-domain-name-example.com",
		Description: "Signup with very long strings",
	},
}

// TestAssertions contains common assertion messages
var TestAssertions = struct {
	UserCreated    string
	UserNotCreated string
	IDGenerated    string
	IDNotGenerated string
	FieldsMatch    string
	FieldsNotMatch string
	NoError        string
	ExpectedError  string
	ResultNotNil   string
	ResultNil      string
	UniqueIDs      string
	NonUniqueIDs   string
}{
	UserCreated:    "User should be created successfully",
	UserNotCreated: "User should not be created",
	IDGenerated:    "User ID should be generated",
	IDNotGenerated: "User ID should not be generated",
	FieldsMatch:    "User fields should match input",
	FieldsNotMatch: "User fields should not match input",
	NoError:        "Operation should not return an error",
	ExpectedError:  "Operation should return an error",
	ResultNotNil:   "Result should not be nil",
	ResultNil:      "Result should be nil",
	UniqueIDs:      "User IDs should be unique",
	NonUniqueIDs:   "User IDs should not be unique",
}
