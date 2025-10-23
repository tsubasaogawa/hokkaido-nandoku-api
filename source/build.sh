#!/bin/bash
set -e

# Create a temporary build directory
rm -rf build
mkdir -p build

# Build the Go application
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o build/bootstrap ./cmd/api

# Copy the data file
cp -r data build/

# Create the zip file
(cd build && zip -r ../hokkaido-nandoku-api.zip .)

# Clean up the build directory
rm -rf build
