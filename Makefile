PACKAGES=$(shell go list ./... | grep -v '/simulation')

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

APP_NAME=litewallet

# TODO: Update the ldflags
ldflags = -X github.com/QOSGroup/litewallet/cmd/version.Name=$(APP_NAME) \
	-X github.com/QOSGroup/litewallet/cmd/version.Version=$(VERSION) \
	-X "github.com/QOSGroup/litewallet/cmd/version.VersionCosmos=github.com/cosmos/cosmos-sdk stargate-3" \
	-X "github.com/QOSGroup/litewallet/cmd/version.VersionEthereum=github.com/ethereum/go-ethereum v1.9.20" \
	-X github.com/QOSGroup/litewallet/cmd/version.Commit=$(COMMIT) 

BUILD_FLAGS := -ldflags '$(ldflags)'

.PHONY: all
all: aar

.PHONY: build
build: go.sum
		go build -mod=readonly $(BUILD_FLAGS) -o ./build/$(APP_NAME) ./cmd/litewallet
		@echo "build ok..."

.PHONY: aar
aar: build
		gomobile bind --target=android -o ./build/$(APP_NAME).aar ./litewallet/mobile/
		@echo "gomobile ok..."

.PHONY: go.sum
go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

# Uncomment when you have some tests
#.PHONY: test
#test:
# 	@go test -mod=readonly $(PACKAGES)

# look into .golangci.yml for enabling / disabling linters
.PHONY: lint
lint:
	@echo "--> Running linter"
	@golangci-lint run
	@go mod verify
