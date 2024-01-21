GO := go
GOBUILD := $(GO) build
GOCLEAN := $(GO) clean

# Name of the executable
BINARY_NAME := certsnap

# Source code files in the current directory
SRC_DIR := .

# Build directory
BUILD_DIR := build

# Target for building the application
build:
	@echo "Building $(BINARY_NAME)..."
	@$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) $(SRC_DIR)
	@echo "Build complete. Executable available in $(BUILD_DIR)/$(BINARY_NAME)"

# Target for cleaning up build artifacts
clean:
	@echo "Cleaning up..."
	@$(GOCLEAN)
	@rm -rf $(BUILD_DIR)

.DEFAULT_GOAL := build
