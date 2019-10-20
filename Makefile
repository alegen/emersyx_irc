emirc.so:
	@go build -buildmode=plugin -o irc.so internal/irc/*

.PHONY: test
test: emirc.so
	@echo "Running the tests with gofmt..."
	@test -z $(shell gofmt -s -l emirc/*.go)
	@echo "Running the tests with go vet..."
	@go vet ./...
	@echo "Running the tests with golint..."
	@golint -set_exit_status $(shell go list ./...)
