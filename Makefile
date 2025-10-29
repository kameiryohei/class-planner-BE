DOCKER_COMPOSE := docker compose
DB_SERVICE := database
DB_NAME := kameiryohei
DB_PORT := 5432
DB_USER := kameiryohei
DB_PASSWORD := kameiryohei

run:
	GO_ENV=dev go run ./main.go

migrate:
	GO_ENV=dev go run migrate/migrate.go

seed:
	GO_ENV=dev go run seed/main.go

compose/up:
	$(DOCKER_COMPOSE) up -d
compose/down:
	$(DOCKER_COMPOSE) down

setup/first:
	make compose/up
	sleep 2
	make migrate

psql:
	$(DOCKER_COMPOSE) exec $(DB_SERVICE) psql -U $(DB_USER) -d $(DB_NAME)

.PHONY: run migrate seed compose/up compose/down setup/first psql
