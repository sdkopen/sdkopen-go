.PHONY: test test-verbose test-cover test-cover-html test-race test-short

## Run all tests
test:
	go test ./...

## Run all tests with verbose output
test-verbose:
	go test ./... -v

## Run tests with coverage report
test-cover:
	go test ./... -cover

## Run tests and generate HTML coverage report
test-cover-html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

## Run tests with race detector
test-race:
	go test ./... -race

## Run only short/fast tests
test-short:
	go test ./... -short
