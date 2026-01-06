# Makefile for backend automation

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run
GOINSTALL=$(GOCMD) install

# Local bin directory
LOCAL_BIN=$(PWD)/bin
export PATH := $(LOCAL_BIN):$(PATH)

# Database parameters
DB_URL="postgresql://user:password@localhost:5433/campus_pilot?sslmode=disable"

# Migration parameters
MIGRATE_PATH=database/migrations
MIGRATE_BIN=$(LOCAL_BIN)/migrate
MIGRATE_CMD=$(MIGRATE_BIN) -database "$(DB_URL)" -path $(MIGRATE_PATH)

.PHONY: all test clean build run migrate-up migrate-down migrate-create migrate-status install-migrate

all: build

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME) $(MIGRATE_BIN)

build:
	$(GOBUILD) -o bin/server ./cmd/server

run:
	$(GORUN) ./cmd/server/main.go

# Database migration commands
migrate-up: $(MIGRATE_BIN)
	@echo "Applying all up migrations..."
	$(MIGRATE_CMD) up

migrate-down: $(MIGRATE_BIN)
	@echo "Reverting last migration..."
	$(MIGRATE_CMD) down 1

migrate-create: $(MIGRATE_BIN)
	@read -p "Enter migration name: " name; \
	$(MIGRATE_CMD) create -ext sql -dir $(MIGRATE_PATH) -seq $$name

migrate-status: $(MIGRATE_BIN)
	@echo "Checking migration status..."
	$(MIGRATE_CMD) version
	
install-migrate:
	@echo "Installing golang-migrate to $(LOCAL_BIN)..."
	mkdir -p $(LOCAL_BIN)
	GOBIN=$(LOCAL_BIN) $(GOINSTALL) github.com/golang-migrate/migrate/v4/cmd/migrate@latest

$(MIGRATE_BIN):
	$(MAKE) install-migrate
