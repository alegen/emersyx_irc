emersyx-irc.so:
	@go build -buildmode=plugin -o emersyx-irc.so internal/irc/*

.PHONY: test
test: emersyx-irc.so
	@echo "Running the tests with gofmt..."
	@test -z $(shell gofmt -s -l internal/irc/*.go)
	@echo "Running the tests with go vet..."
	@go vet ./...
	@echo "Running the tests with golint..."
	@golint -set_exit_status $(shell go list ./...)
