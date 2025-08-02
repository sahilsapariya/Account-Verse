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

> if you don't have make, you can cd into server dir and build using the go build command. In that case you will have to build docs manually using npm run build on both dirs.

### Setup

#### air.toml - File
```bash
#For windows

root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main.exe"
  cmd = "go build -o ./tmp/main.exe ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false

```

```bash
#For mac/ubuntu

root = "."
testdata_dir = "testdata"
tmp_dir = "tmp"

[build]
  args_bin = []
  bin = "./tmp/main"
  cmd = "go build -o ./tmp/main ."
  delay = 1000
  exclude_dir = ["assets", "tmp", "vendor", "testdata"]
  exclude_file = []
  exclude_regex = ["_test.go"]
  exclude_unchanged = false
  follow_symlink = false
  full_bin = ""
  include_dir = []
  include_ext = ["go", "tpl", "tmpl", "html"]
  kill_delay = "0s"
  log = "build-errors.log"
  send_interrupt = false
  stop_on_root = false

[color]
  app = ""
  build = "yellow"
  main = "magenta"
  runner = "green"
  watcher = "cyan"

[log]
  time = false

[misc]
  clean_on_exit = false
``` 

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

