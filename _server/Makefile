BUILD_FLAGS := -ldflags "-s -w"
OUTPUT_PATH := ./bin/fsserver

.PHONY: default

default: build

build:
	go build $(BUILD_FLAGS) -o $(OUTPUT_PATH)
	@echo "Build done"
