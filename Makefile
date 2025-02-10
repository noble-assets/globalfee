.PHONY: proto-format proto-lint proto-gen license format lint test-unit build local-image test-e2e
all: proto-all format lint test-unit build local-image test-e2e

###############################################################################
###                                  Build                                  ###
###############################################################################

build:
	@echo "🤖 Building simd..."
	@cd simapp && make build 1> /dev/null
	@echo "✅ Completed build!"

###############################################################################
###                                 Tooling                                 ###
###############################################################################

gofumpt_cmd=mvdan.cc/gofumpt
golangci_lint_cmd=github.com/golangci/golangci-lint/cmd/golangci-lint

FILES := $(shell find $(shell go list -f '{{.Dir}}' ./...) -name "*.go" -a -not -name "*.pb.go" -a -not -name "*.pb.gw.go" -a -not -name "*.pulsar.go" | sed "s|$(shell pwd)/||g")
license-update:
	@go-license --config .github/license.yml --remove $(FILES)
	@go-license --config .github/license.yml $(FILES)

license:
	@go-license --config .github/license.yml $(FILES)

format:
	@echo "🤖 Running formatter..."
	@go run $(gofumpt_cmd) -l -w .
	@echo "✅ Completed formatting!"

lint:
	@echo "🤖 Running linter..."
	@go run $(golangci_lint_cmd) run --timeout=10m
	@echo "✅ Completed linting!"

###############################################################################
###                                Protobuf                                 ###
###############################################################################

BUF_VERSION=1.42
BUILDER_VERSION=0.15.1

proto-all: proto-format proto-lint proto-gen

proto-format:
	@echo "🤖 Running protobuf formatter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) format --diff --write
	@echo "✅ Completed protobuf formatting!"

proto-gen:
	@echo "🤖 Generating code from protobuf..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		ghcr.io/cosmos/proto-builder:$(BUILDER_VERSION) sh ./proto/generate.sh
	@echo "✅ Completed code generation!"

proto-lint:
	@echo "🤖 Running protobuf linter..."
	@docker run --rm --volume "$(PWD)":/workspace --workdir /workspace \
		bufbuild/buf:$(BUF_VERSION) lint
	@echo "✅ Completed protobuf linting!"

###############################################################################
###                                 Testing                                 ###
###############################################################################

local-image:
	@echo "🤖 Building image..."
	@heighliner build --file ./e2e/chains.yaml --chain noble-globalfee-simd --local
	@echo "✅ Completed build!"

test-unit:
	@echo "🤖 Running unit tests..."
	@go test -cover -coverprofile=coverage.out -race -v ./keeper/...
	@echo "✅ Completed unit tests!"

test-e2e:
	@echo "🤖 Running e2e tests..."
	@cd e2e && go test -timeout 15m -race -v ./...
	@echo "✅ Completed e2e tests!"
