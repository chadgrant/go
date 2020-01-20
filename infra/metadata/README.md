# Package metadata

This package is intended to receive build flags (ldflags). Bake that information into the binary and expose it on a /metadata endpoint. The metadata will also be included in docker build labels


## Usage
```go
	md := metadata.NewHandler()
	http.HandleFunc("/metadata", md.Metadata)
```

## Makefile Sample Variables

```make
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

```


## Makefile Build Target Sample

```make

PKG=github.com/shipt/shipt-tofu/metadata
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

build: ## Builds go binary
	go build -ldflags $(LDFLAGS)

```

## GET /metadata response

Returns the metadata of the service. Think of it as running --version on a local binary, just on a webservice!

> GET /metadata

| Property           | Description                                              |     Example                                |
| -------------------|----------------------------------------------------------|--------------------------------------------|
| vendor             | name of the vendor                                       | shipt                                      |
| group              | name of the team responsible for the service             | member                                     |
| service            | name of the service, normalized with no spaces           | sample-api                                 |
| friendly           | friendly name of the service                             | Sample API                                 |
| description        | optional short description of service                    | Sample API                                 |
| repo               | optional url of service repo                             | https://github.com/shipt/sample-api        |
| build_number       | The build number                                         | 2.5.7                                      |
| built_by           | The user that did the build                              | cgrant                                     |
| build_time         | When the build was done                                  | 2015-03-12T19:40:18.877Z                   |
| git_hash           | The git sha1 hash of the build                           | d567d2650318f704747204815adedd2396a203f5   |
| git_branch         | The git branch of the build                              | master                                     |
| compiler_version   | The compiler version                                     | 1.5                                        |
| version            | Current version of the metadata spec                     | 1                                          |
