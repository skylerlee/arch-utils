BUILD_FLAGS := -ldflags "-s -w"
OUTPUT_PATH := $(PWD)/bin

.PHONY: default

default: build

build: export GOBIN = $(OUTPUT_PATH)
build:
	go install $(BUILD_FLAGS) .
	@echo "Build done"
