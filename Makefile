# Go tools for cross-platform support
MKDIR := go run ./mkdir/main.go
RM    := go run ./clean/main.go

ifeq ($(OS),Windows_NT)
	EXE := .exe
else
	EXE :=
endif

.PHONY: help build test clean install mod-tidy

.DEFAULT_GOAL := help

help:
	@echo "gonest-tools commands:"
	@echo ""
	@echo "  make build      # Build all tools to ./bin/"
	@echo "  make test       # Test all tools"
	@echo "  make install    # Install all tools to GOPATH/bin"
	@echo "  make clean      # Clean build artifacts"

build: build-badge build-clean build-tag build-mkdir
	@echo "All tools built in ./bin/"

build-badge:
	@$(MKDIR) bin
	@cd badge && go build -o ../bin/gonest-badge$(EXE)
	@echo "gonest-badge built"

build-clean:
	@$(MKDIR) bin
	@cd clean && go build -o ../bin/gonest-clean$(EXE)
	@echo "gonest-clean built"

build-tag:
	@$(MKDIR) bin
	@cd tag && go build -o ../bin/gonest-tag$(EXE)
	@echo "gonest-tag built"

build-mkdir:
	@$(MKDIR) bin
	@cd mkdir && go build -o ../bin/gonest-mkdir$(EXE)
	@echo "gonest-mkdir built"

test:
	@cd tag && go test -v ./...
	@echo "All tests passed"

install:
	@cd badge && go install
	@cd clean && go install
	@cd tag && go install
	@cd mkdir && go install
	@echo "Tools installed"

clean:
	@$(RM) bin
	@echo "Cleaned"

mod-tidy:
	@cd badge && go mod tidy
	@cd clean && go mod tidy
	@cd tag && go mod tidy
	@cd mkdir && go mod tidy
	@echo "All modules tidied"
