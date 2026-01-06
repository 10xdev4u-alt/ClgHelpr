# Makefile for backend automation

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

# Database parameters
DB_URL="postgresql://user:password@localhost:5433/campus_pilot?sslmode=disable"

# Migration parameters
MIGRATE_PATH=database/migrations
MIGRATE_CMD=migrate -database "$(DB_URL)" -path $(MIGRATE_PATH)

.PHONY: all test clean build run

all: build

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

build:
	$(GOBUILD) -o bin/server ./cmd/server

run:
	$(GORUN) ./cmd/server/main.go

# Database migration commands
migrate-up:
	@echo "Applying all up migrations..."
	$(MIGRATE_CMD) up

migrate-down:
	@echo "Reverting last migration..."
	$(MIGRATE_CMD) down 1

migrate-create:
	@read -p "Enter migration name: " name; \
	$(MIGRATE_CMD) create -ext sql -dir $(MIGRATE_PATH) -seq $$name

migrate-status:
	@echo "Checking migration status..."
	$(MIGRATE_CMD) version
	
install-migrate:
	@echo "Installing golang-migrate..."
	$(GOGET) -u github.com/golang-migrate/migrate/v4/cmd/migrate
