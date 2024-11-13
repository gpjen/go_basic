
.PHONY: run migrate migrate-up migrate-down migrate-status build test

# Default environment
ENV ?= development

# Application commands
run:
	@go run cmd/app/main.go

build:
	@go build -o bin/app cmd/app/main.go

# Migration commands
migrate-status:
	@./scripts/migrate.sh -e $(ENV) -s

migrate-up:
	@./scripts/migrate.sh -e $(ENV) -u

migrate-down:
	@./scripts/migrate.sh -e $(ENV) -d

# Environment specific migrations
migrate-prod-up:
	@./scripts/migrate.sh -e production -u

migrate-prod-down:
	@./scripts/migrate.sh -e production -d

migrate-staging-up:
	@./scripts/migrate.sh -e staging -u

migrate-staging-down:
	@./scripts/migrate.sh -e staging -d

# Test
test:
	@go test ./... -v

# Add necessary permissions to scripts
init:
	@chmod +x scripts/migrate.sh