# Test Suite Documentation

This directory contains comprehensive test cases for the Account-Verse server application.

## Test Structure

### Files Overview

- **`signup_test.go`** - Main test file for user signup functionality
- **`test_helpers.go`** - Utility functions and helpers for testing
- **`test_config.go`** - Test configuration and common test scenarios
- **`health_check_test.go`** - Basic health check endpoint test

### Test Categories

#### 1. Unit Tests
- **`TestSignupResolver_*`** - Tests for the signup resolver function
- **`TestAddUser_*`** - Tests for database user creation operations

#### 2. Integration Tests
- **`TestSignupFlow_Integration`** - End-to-end signup flow testing

#### 3. Performance Tests
- **`BenchmarkSignupResolver`** - Performance benchmarks for signup operations
- **`BenchmarkAddUser`** - Performance benchmarks for database operations

#### 4. Edge Case Tests
- **`TestAddUser_SpecialCharacters`** - Unicode and special character handling
- **`TestAddUser_LongStrings`** - Long string handling
- **`TestAddUser_Concurrent`** - Multiple user creation testing

## Running Tests

### Run All Tests
```bash
make test
```

### Run Specific Test Files
```bash
go test ./test/signup_test.go ./test/test_helpers.go -v
```

### Run Tests with Coverage
```bash
make test-coverage
```

### Run Benchmark Tests
```bash
go test ./test/... -bench=. -benchmem
```

### Run Tests with Race Detection
```bash
make test-race
```

## Test Coverage

The current test suite covers:

### ✅ Covered Functionality
- [x] User creation via signup resolver
- [x] Database user insertion
- [x] UUID generation for users
- [x] Basic input validation (empty fields)
- [x] Special character handling
- [x] Long string handling
- [x] Multiple user creation
- [x] Error handling for database operations
- [x] Performance benchmarking

### 🔄 Future Test Coverage (for next PR)
- [ ] Email validation
- [ ] Password hashing
- [ ] JWT token generation
- [ ] Duplicate email handling
- [ ] Input sanitization
- [ ] Rate limiting
- [ ] Authentication middleware
- [ ] Error response formatting
- [ ] Database transaction handling
- [ ] Integration with external services

## Test Data

### Test Users
The test suite uses predefined test data from `TestUserData`:
- **ValidUser1**: "John Doe" <john.doe@example.com>
- **ValidUser2**: "Jane Smith" <jane.smith@example.com>
- **SpecialCharUser**: "José María García-López" <jose.maria+test@example.com>
- **LongStringUser**: Very long name and email for boundary testing

### Test Scenarios
Predefined scenarios in `TestScenarios`:
- **ValidSignup**: Normal valid signup
- **InvalidEmail**: Invalid email format
- **EmptyFields**: Empty name and email
- **SpecialChars**: Unicode and special characters
- **LongStrings**: Boundary testing with long strings

## Database Testing

### Test Database Setup
- Uses in-memory SQLite database (`:memory:`)
- Isolated test environment for each test
- Automatic cleanup after each test
- Thread-safe database initialization

### Test Database Configuration
```go
// Default test configuration
DBType: "sqlite"
DBName: ":memory:"
LogLevel: "error"
```

## Performance Benchmarks

Current benchmark results (on Intel i5-7300U @ 2.60GHz):
- **SignupResolver**: ~78,328 ns/op (78.3 µs per operation)
- **AddUser**: ~65,355 ns/op (65.4 µs per operation)

## Test Helpers

### Database Helpers
- `SetupTestDB(t)` - Initialize test database
- `TeardownTestDB()` - Clean up test database
- `AddUserToTestDB(t, user)` - Add user to test database

### Data Helpers
- `CreateTestUser(name, email)` - Create test user struct
- `CreateTestUserWithID(id, name, email)` - Create user with specific ID
- `CreateTestSignUpInput(name, email)` - Create signup input struct

### Environment Helpers
- `SetupTestEnvironment(t, config)` - Set up test environment
- `DefaultTestConfig()` - Get default test configuration

## Best Practices

### Test Naming
- Use descriptive test names: `TestSignupResolver_Success`
- Group related tests: `TestAddUser_*`
- Use subtests for scenarios: `TestSignupFlow_Integration`

### Test Structure
1. **Setup** - Initialize test database and data
2. **Execute** - Run the function being tested
3. **Assert** - Verify expected outcomes
4. **Cleanup** - Clean up resources (handled by defer)

### Error Handling
- Always check for unexpected errors
- Use descriptive assertion messages
- Test both success and failure cases

## Continuous Integration

The test suite is integrated with:
- **Pre-commit hooks** - Run tests before commits
- **Make targets** - Easy test execution
- **Coverage reporting** - Track test coverage
- **Benchmark tracking** - Monitor performance

## Adding New Tests

When adding new tests:
1. Follow the existing naming conventions
2. Use the helper functions for setup/teardown
3. Add test data to `TestUserData` if needed
4. Update this documentation
5. Ensure tests are isolated and repeatable

## Debugging Tests

### Verbose Output
```bash
go test ./test/... -v
```

### Run Single Test
```bash
go test ./test/... -run TestSignupResolver_Success -v
```

### Debug with Print Statements
```go
t.Logf("Debug info: %+v", variable)
```

## Test Environment Variables

- `DB_TYPE` - Database type (default: "sqlite")
- `DB_NAME` - Database name (default: ":memory:")
- `LOG_LEVEL` - Log level (default: "error" for tests) 