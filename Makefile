include .env

YELLOW := \033[0;33m
NC := \033[0m

MIGRATIONS_PATH := database/migrations

migrate:
	@echo "${YELLOW}===> Enter the name of the table to create : ${NC}"; \
	read table; \
	migrate create -ext sql -dir $(MIGRATIONS_PATH) -seq $$table

db-up:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

db-down:
	migrate -path $(MIGRATIONS_PATH) -database "$(DB_TYPE)://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

run-local:
	docker-compose up -d
	db-up