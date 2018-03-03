emirc.so: goget
	go build -buildmode=plugin -o emirc.so emirc/*

.PHONY: goget
goget:
	go get emersyx.net/emersyx_apis/emcomapi
	go get emersyx.net/emersyx_apis/emircapi
	go get github.com/fluffle/goirc/client
	go get github.com/fluffle/goirc/logging
	go get github.com/fluffle/goirc/state

.PHONY: test
test: emirc.so
	@echo "Running the tests with gofmt, go vet and golint..."
	@test -z $(shell gofmt -s -l emirc/*.go)
	@go vet ./...
	@golint -set_exit_status $(shell go list ./...)
