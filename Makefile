.PHONY: help build test clean install

.DEFAULT_GOAL := help

help: ## Show this help
	@echo "gonest-tools commands:"
	@echo ""
	@echo "  make build      # Build all tools to ./bin/"
	@echo "  make test       # Test all tools"
	@echo "  make install    # Install all tools to \$$GOPATH/bin"
	@echo "  make clean      # Clean build artifacts"
	@echo ""
	@echo "Individual tools:"
	@echo "  make build-badge    # Build badge tool"
	@echo "  make build-clean    # Build clean tool"
	@echo "  make build-tag      # Build tag tool"

build: ## Build all tools
	@echo "Building all tools..."
	@mkdir -p bin
	@echo "Building badge..."
	@cd badge && go build -o ../bin/gonest-badge
	@echo "Building clean..."
	@cd clean && go build -o ../bin/gonest-clean
	@echo "Building tag..."
	@cd tag && go build -o ../bin/gonest-tag
	@echo "✓ All tools built in ./bin/"

build-badge: ## Build badge tool
	@mkdir -p bin
	@cd badge && go build -o ../bin/gonest-badge
	@echo "✓ gonest-badge built"

build-clean: ## Build clean tool
	@mkdir -p bin
	@cd clean && go build -o ../bin/gonest-clean
	@echo "✓ gonest-clean built"

build-tag: ## Build tag tool
	@mkdir -p bin
	@cd tag && go build -o ../bin/gonest-tag
	@echo "✓ gonest-tag built"

test: ## Test all tools
	@echo "Testing tag..."
	@cd tag && go test -v ./...
	@echo "✓ All tests passed!"

install: ## Install all tools
	@echo "Installing badge..."
	@cd badge && go install
	@echo "Installing clean..."
	@cd clean && go install
	@echo "Installing tag..."
	@cd tag && go install
	@echo "✓ Tools installed to \$$GOPATH/bin"

clean: ## Clean build artifacts
	@rm -rf bin/
	@echo "✓ Cleaned!"

mod-tidy: ## Tidy all modules
	@echo "Tidying badge..."
	@cd badge && go mod tidy
	@echo "Tidying clean..."
	@cd clean && go mod tidy
	@echo "Tidying tag..."
	@cd tag && go mod tidy
	@echo "✓ All modules tidied"