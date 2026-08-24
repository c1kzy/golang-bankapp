include .env

export

export PROJECT_ROOT=$(CURDIR)

bankapp-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/main.go

env-up:
	@docker compose up -d bankapp-postgres

env-down:
	@docker compose down bankapp-postgres

env-cleanup:
	@read -p "Clean all volume files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down bankapp-postgres && \
		rm -rf ${PROJECT_ROOT}/out/pgdata && \
		echo "Env files cleared"; \
	else \
		echo "Cleanup canceled"; \
	fi

logs-cleanup:
	@read -p "Clean all log files? [y/N]: " ans; \
	if [ "$$ans" = "y" ]; then \
		rm -rf ${PROJECT_ROOT}/out/logs && \
		echo "Log files cleared"; \
	else \
		echo "Log cleanup canceled"; \
	fi

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "seq param is missing. Example: make migrate-create seq=init"; \
		exit 1; \
	fi; \
	docker compose run --rm bankapp-postgres-migrate \
		create \
		-ext sql \
		-dir migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "action param is missing. Example: make migrate-action action=up"; \
		exit 1; \
	fi; \
	docker compose run --rm bankapp-postgres-migrate \
		-path migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@bankapp-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"