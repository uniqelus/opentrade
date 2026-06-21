BUILD_TARGETS := $(shell ls cmd)
BUILD_GOOS ?= $(shell go env GOOS)
BUILD_GOARCH ?= $(shell go env GOARCH)
BUILD_CGO_ENABLED ?= 0
BUILD_FLAGS ?= -ldflags="-s -w"

.PHONY: build
build: $(addsuffix -build, $(BUILD_TARGETS))
%-build:
	CGO_ENABLED=$(BUILD_CGO_ENABLED) GOOS=$(BUILD_GOOS) GOARCH=$(BUILD_GOARCH) \
	go build $(BUILD_FLAGS) -o bin/$* ./cmd/$*/main.go

.PHONY: lint
lint:
	golangci-lint run $(if $(LINT_FIX),,--fix) ./...

.PHONY: image
image: $(addsuffix -image, $(BUILD_TARGETS))
%-image: build
	docker build -f docker/$*.dockerfile -t opentrade/$*:latest .
