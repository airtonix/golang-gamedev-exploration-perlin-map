# JUSTFILE
# https://github.com/casey/just
#

default:
    @just --choose

help:
    @just --list

# Project setup
setup:
    @echo '⛳ Begin project setup'
    go get


# Repo cleanup
fix:
    go mod tidy

# Linting
lint:
    golangci-lint run 

# Creates a snapshot build
build:
    goreleaser build --snapshot --rm-dist

# Creates a release build
release:
    go clean
    go generate ./...
    go mod tidy
    goreleaser build

# Test out game inlocal wasm server
serve:
    @echo '🍱 WASM being served on http://localhost:8080'
    @wasmserve .
