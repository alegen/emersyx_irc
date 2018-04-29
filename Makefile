emirc.so: goget
	@go build -buildmode=plugin -o emirc.so emirc/*

.PHONY: goget
goget:
	@go get emersyx.net/emersyx/api
	@go get github.com/fluffle/goirc/client
	@go get github.com/fluffle/goirc/logging
	@go get github.com/fluffle/goirc/state
	@go get github.com/golang/lint/golint
	@go get github.com/BurntSushi/toml

.PHONY: test
test: emirc.so
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emirc/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
