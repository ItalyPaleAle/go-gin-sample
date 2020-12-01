#!/bin/sh

# Create directories
mkdir -p .bin .out

# Build the frontend app
(cd frontend && npm ci && npm run build)

# Fetch pkger if needed
test -f .bin/pkger || \
	  curl -sf https://gobinaries.com/github.com/markbates/pkger/cmd/pkger@v0.17.1 | PREFIX=.bin/ sh

# Run pkger
.bin/pkger

# Build the app
go build -o .out/app .
