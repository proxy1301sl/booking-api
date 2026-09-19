include .env
export

export PROJECT_ROOT := $(shell pwd)

env-up:
	docker compose up -d postgres

env-down:
	docker compose down

env-cleanup:
	docker compose down postgres && rm -rf out/pgdata


seq ?= migration

migrate-create:
	docker compose run --rm migrate \
        		create -ext sql -dir /migrations -seq "$(seq)"




action ?= version

.PHONY: migrate-up migrate-down migrate-version migrate-action

migrate-up:
	@$(MAKE) migrate-action action=up

migrate-down:
	@$(MAKE) migrate-action action=down

migrate-version:
	@$(MAKE) migrate-action action=version

migrate-action:
	docker compose run --rm migrate \
		-path /migrations \
		-database "postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@postgres:5432/${POSTGRES_DB}?sslmode=disable" \
		"$(action)"





