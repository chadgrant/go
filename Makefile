.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Builds go binary
	go build

.PHONY: test
test: ## Run Tests
	CGO_ENABLED=1 go test -v -race ./...

.PHONY: test-integration
test-integration: # Run integration tests
	CGO_ENABLED=1 TEST_INTEGRATION=1 go test -race -v ./...
	
clean: ## Cleans directory of temp files
	-@GO111MODULE=off go clean -i
	-@rm -f $(OUT_DIR)$(BINARY_NAME)
	-@rm -f coverage.html profile.out cpu.prof coverage.txt

tidy: ## Run goimports and go fmt on all *.go files
ifeq (,$(shell type goimports 2>/dev/null))
	go get golang.org/x/tools/cmd/goimports
endif
	@go fmt ./...
	@goimports -w $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)

lint: ## Execute linter
ifeq (,$(shell type golangci-lint 2>/dev/null))
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(shell go env GOPATH)/bin v1.22.2
endif
	golangci-lint run --timeout=300s --skip-dirs-use-default --exclude="should have comment or be unexported"  ./...

reportcard: ## Display go report card status
ifeq (,$(shell type gometalinter 2>/dev/null))
	## has not transitioned to golangci-lint yet
	cd $(GOPATH); curl -L https://git.io/vp6lP | sh
endif
ifeq (,$(shell type goreportcard-cli 2>/dev/null))
	go get github.com/gojp/goreportcard/cmd/goreportcard-cli
endif
	goreportcard-cli

cover: ## Run go code coverage tool
	@go test -covermode=atomic -coverpkg=./... -coverprofile=profile.out ./...
	@go tool cover -func=profile.out
	@go tool cover -html=profile.out -o coverage.html