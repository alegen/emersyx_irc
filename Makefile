.PHONY: emirc test

emirc:
	go get -t -buildmode=plugin ./emirc

test:
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emirc/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
