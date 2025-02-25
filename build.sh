#!/bin/bash

# Mac only for now

# Move to the app directory first
cd app

# Create directory for the PHP binary relative to build.go
mkdir -p php_binary

# Extract PHP binary from zip (adjusting path since we're in go directory)
unzip ../vendor/nativephp/php-bin/bin/mac/arm64/php-8.4.zip -d php_binary

# See what we got
ls -la php_binary

# Build the Go binary
GOOS=darwin GOARCH=arm64 go build -o ../build/mac-arm64 wrapper.go

# Clean up
rm -rf php_binary

# Go back to original directory
cd ..
