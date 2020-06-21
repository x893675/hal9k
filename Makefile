GOPATH ?= $(shell go env GOPATH)

apps = 'account'

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')
REPO_URL ?= $(shell git ls-remote --get-url origin)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_REF ?= $(shell git rev-parse --verify HEAD)

.PHONY: build
build:
	CGO_ENABLED=0 go build -o dist/$$app -a -ldflags "-w -s -X hal9k/pkg/version.Version=${VERSION}" ./cmd/
