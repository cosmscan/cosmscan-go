GOBIN = ./bin
GORUN = env GO111MODULE=on go run
GOBUILD = env GO111MODULE=on go build

PACKAGES_UNIT=$(shell go list ./...)

build:
	if [ ! -d $(GOBIN) ]; then mkdir $(GOBIN); fi
	$(GOBUILD) -v -o $(GOBIN) ./...

test:
	@go test -mod=readonly $(PACKAGES_UNIT)

.PHONEY: build test