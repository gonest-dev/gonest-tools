ifeq ($(OS),Windows_NT)
	EXE := .exe
	# PowerShell é mais consistente no Windows do que comandos CMD
	MKDIR := powershell -NoProfile -Command "if (!(Test-Path bin)) { New-Item -Path bin -ItemType Directory }"
	RM := powershell -NoProfile -Command "if (Test-Path bin) { Remove-Item -Recurse -Force bin }"
else
	EXE :=
	MKDIR := mkdir -p bin
	RM := rm -rf bin
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

build: build-badge build-clean build-tag
	@echo "All tools built in ./bin/"

build-badge:
	@$(MKDIR)
	@cd badge && go build -o ../bin/gonest-badge$(EXE)
	@echo "gonest-badge built"

build-clean:
	@$(MKDIR)
	@cd clean && go build -o ../bin/gonest-clean$(EXE)
	@echo "gonest-clean built"

build-tag:
	@$(MKDIR)
	@cd tag && go build -o ../bin/gonest-tag$(EXE)
	@echo "gonest-tag built"

test:
	@cd tag && go test -v ./...
	@echo "All tests passed"

install:
	@cd badge && go install
	@cd clean && go install
	@cd tag && go install
	@echo "Tools installed"

clean:
	@$(RM)
	@echo "Cleaned"

mod-tidy:
	@cd badge && go mod tidy
	@cd clean && go mod tidy
	@cd tag && go mod tidy
	@echo "All modules tidied"
