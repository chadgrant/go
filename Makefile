VENDOR?=shipt
GROUP?=member

SERVICE?=name_of_service
SERVICE_FRIENDLY?=Friendly Name of Service
SERVICE_DESCRIPTION?=Sample service description
SERVICE_URL?=http://localhost/url/of/deployment

BUILD_REPO?=https://github.com/shipt/sample-service
BUILD_NUMBER?=$(subst v,,$(shell git describe --tags --dirty --match=v* 2> /dev/null || echo 1.0.0))
BUILD_BRANCH?=$(shell git rev-parse --abbrev-ref HEAD)
BUILD_HASH?=$(shell git rev-parse HEAD)
BUILD_DATE?=$(shell TZ=UTC date -u '+%Y-%m-%dT%H:%M:%SZ')

ifdef BUILD_HASH
 ifndef BUILD_USER
	BUILD_USER?=$(shell git --no-pager show -s --format='%ae' $(BUILD_HASH) 2> /dev/null || echo $(USER))
 endif
else
 BUILD_USER?=$(USER)
endif

PKG=github.com/chadgrant/infra/metadata
LDFLAGS="-w -s \
		-X '$(PKG).Vendor=$(VENDOR)' \
		-X '$(PKG).Group=$(GROUP)' \
		-X '$(PKG).Service=$(SERVICE)' \
		-X '$(PKG).Friendly=$(SERVICE_FRIENDLY)' \
		-X '$(PKG).Description=$(SERVICE_DESCRIPTION)' \
		-X '$(PKG).Repo=$(BUILD_REPO)' \
		-X '${PKG}.BuildNumber=$(BUILD_NUMBER)' \
		-X '$(PKG).BuiltBy=$(BUILD_USER)' \
		-X '$(PKG).BuildTime=$(BUILD_DATE)' \
		-X '$(PKG).GitHash=$(BUILD_HASH)' \
		-X '$(PKG).GitBranch=$(BUILD_BRANCH)' \
		-X '$(PKG).CompilerVersion=$(shell go version)'"

.DEFAULT_GOAL := help
.EXPORT_ALL_VARIABLES:

.PHONY: help
help:
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: generate ## Builds go binary
	go build -ldflags $(LDFLAGS)

run: generate ## Runs the main.go file
	go run -ldflags $(LDFLAGS) main.go

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
	@if ! type "goimports" > /dev/null 2>&1; then \
		go get -u golang.org/x/tools/cmd/goimports; \
	fi
	@go fmt ./...
	@goimports -w $(shell go list -f {{.Dir}} ./... | grep -v /vendor/)

lint: ## Execute linter
	@if ! type "golangci-lint" > /dev/null 2>&1; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
		| sh -s -- -b $(shell go env GOPATH)/bin v1.22.2; \
	fi
	golangci-lint run --timeout=300s --skip-dirs-use-default --exclude="should have comment or be unexported"  ./...

reportcard: ## Display go report card status
	@if ! type "gometalinter" > /dev/null 2>&1; then \
		cd $(GOPATH); curl -L https://git.io/vp6lP | sh; \
	fi 
	@if ! type "goreportcard-cli" > /dev/null 2>&1; then \
		go get github.com/gojp/goreportcard/cmd/goreportcard-cli; \
	fi
	goreportcard-cli

cover: ## Run go code coverage tool
	@go test -covermode=atomic -coverpkg=./... -coverprofile=profile.out ./...
	@go tool cover -func=profile.out
	@go tool cover -html=profile.out -o coverage.html

generate: embed ## Generated go code

embed: ## Embed static resources (go-bindata)
	@if ! type "go-bindata" > /dev/null 2>&1; then \
		go get -u github.com/go-bindata/go-bindata/...; \
	fi
	cd ./infra/schema; go-bindata -pkg schema *.json
	cd ./infra/health; go-bindata -pkg health *.json
	cd ./infra/metadata; go-bindata -pkg metadata *.json

publish-schema: ## Copies schema to S3
	aws s3 cp ./infra/health/schema.json s3://schemas.sentex.io/service/health.json
	aws s3 cp ./infra/metadata/schema.json s3://schemas.sentex.io/service/metadata.json
	aws s3 cp ./infra/schema/schema.json s3://schemas.sentex.io/service/schemalist.json