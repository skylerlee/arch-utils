BUILD_FLAGS := -ldflags "-s -w"
OUTPUT_PATH := ./bin/gotif

.PHONY: default

default: build

build:
	go build $(BUILD_FLAGS) -o $(OUTPUT_PATH)
	@echo "Build done"
