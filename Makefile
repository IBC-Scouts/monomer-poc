# Makefile for building the peptide execution engine.

# Target directory for the build output
BUILD_DIR := build

.PHONY: build-peptide

build-peptide:
	@echo "Building peptide..."
	@go build -o $(BUILD_DIR)/peptide cmd/peptide/main.go
	@echo "Build complete!"