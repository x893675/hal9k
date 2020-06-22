GOPATH ?= $(shell go env GOPATH)

apps = 'hal9k'

VERSION ?= $(shell git rev-parse --short HEAD)-$(shell date -u '+%Y%m%d%I%M%S')
REPO_URL ?= $(shell git ls-remote --get-url origin)
BRANCH ?= $(shell git rev-parse --abbrev-ref HEAD)
COMMIT_REF ?= $(shell git rev-parse --verify HEAD)

.PHONY: build
build:
	CGO_ENABLED=0 go build -o dist/$$app -a -ldflags "-w -s -X hal9k/pkg/version.Version=${VERSION}" ./cmd/


.PHONY: image
image:
	docker build --build-arg REPO_URL=$(REPO_URL) --build-arg BRANCH=$(BRANCH) --build-arg COMMIT_REF=$(COMMIT_REF) --build-arg VERSION=$(VERSION) -f Dockerfile -t hal9k:latest .