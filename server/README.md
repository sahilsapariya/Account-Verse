# Account-Verse Server

A GraphQL server built with Go and Gin framework.

## Features

- GraphQL API with gqlgen
- Hot reloading with Air
- Database support (SQLite, MongoDB)
- Comprehensive testing
- Linting and formatting
- Docker support

## Development

### Prerequisites

- Go 1.23+
- Make

### Setup

```bash
# Install development tools
make install-tools

# Setup development environment
make setup

# Run in development mode
make dev
```

### Available Commands

Run `make help` to see all available commands.

### Testing

```bash
# Run tests
make test

# Run tests with coverage
make test-coverage
```

### Building

```bash
# Build for current platform
make build

# Build for all platforms
make build-all
```

## API

The GraphQL playground is available at `http://localhost:8080/` when the server is running.

