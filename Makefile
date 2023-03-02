VERSION=$(shell ./tools/image_tag | cut -d, -f 1)

GIT_REVISION := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

GO_OPT= -ldflags "-X main.Branch=$(GIT_BRANCH) -X main.Revision=$(GIT_REVISION) -X main.Version=$(VERSION)"
GOBIN = ./bin
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

PACKAGES_UNIT=$(shell go list ./...)
FILES_TO_FMT=$(shell find . -type f -name "*.go" -not -name "*.pb.go" -not -name "*pb.gw.go")
DOCKER_LINT_IMAGE ?= golangci/golangci-lint:v1.51.1

### lint
.PHONY: lint
lint:
	docker run --rm -v $(pwd):/app -w /app $(DOCKER_LINT_IMAGE) golangci-lint run -v

### fmt
.PHONY: fmt check-fmt
fmt:
	@gofmt -s -w $(FILES_TO_FMT)
	@goimports -w $(FILES_TO_FMT)

build:
	go build $(GO_OPT) -o ./bin/$(GOOS)/cosmscan-$(GOARCH) ./cmd/cosmscan

test:
	@go test -mod=readonly $(PACKAGES_UNIT)

test-with-cover:
	@go test -race -timeout 30m -count=1 -cover $(PACKAGES_UNIT)

.PHONEY: build test