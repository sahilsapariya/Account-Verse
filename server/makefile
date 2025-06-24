.PHONY: gqlgen-generate test test-coverage

gqlgen-generate:
	go get github.com/99designs/gqlgen/codegen/config@v0.17.75
	go get github.com/99designs/gqlgen/internal/imports@v0.17.75
	go get github.com/99designs/gqlgen/api@v0.17.75
	go get github.com/99designs/gqlgen@v0.17.75
	go run github.com/99designs/gqlgen generate

test:
	go test ./test/...

test-coverage:
	go test ./test/... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html
