GOBIN = ./bin
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

PACKAGES_UNIT=$(shell go list ./...)

build:
	if [ ! -d $(GOBIN) ]; then mkdir $(GOBIN); fi
	$(GOBUILD) -v -o $(GOBIN) ./...

test:
	@go test -mod=readonly $(PACKAGES_UNIT)

run-indexer:
	$(GORUN) ./cmd/cosmscan/main.go --config-file ./config.yml

run-server:
	$(GORUN) ./cmd/server/main.go --config-file ./config-server.yml

.PHONEY: build test run-indexer run-server