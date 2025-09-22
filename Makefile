# Makefile for Task-Ninja

BINARY_NAME=task-ninja
BUILD_DIR=build

.PHONY: all build clean run

all: build

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) main.go

run: build
	./$(BUILD_DIR)/$(BINARY_NAME)

clean:
	rm -rf $(BUILD_DIR)
