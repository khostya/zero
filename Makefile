DEFAULT_PG_URL=postgres://user:password@localhost:5432/postgres?sslmode=disable
COMPOSE=docker compose -f docker-compose.yml --env-file ./build/docker.env
COMPOSE_APP=${COMPOSE} -p app
COMPOSE_POSTGRES=${COMPOSE} -p postgres
COMPOSE_DEV=${COMPOSE} -p dev
LOCAL_BIN:=$(CURDIR)/bin

.PHONY: .up
up:
	${COMPOSE} up -d --build

.PHONY: .down
down:
	${COMPOSE} down

.PHONY: up-postgres
up-postgres:
	${COMPOSE_POSTGRES} up -d --build postgres migrate

.PHONY: down-postgres
down-postgres:
	${COMPOSE_POSTGRES} down

.PHONY: up-dev
up-dev:
	${COMPOSE_DEV} up -d --build postgres migrate-server service docs

.PHONY: down-dev
down-dev:
	${COMPOSE_DEV} down

.PHONY: migration-up
migration-up:
	$(eval PG_URL?=$(DEFAULT_PG_URL))
	$(LOCAL_BIN)/goose -dir ./migrations postgres "$(PG_URL)" up

.PHONY: migration-down
migration-down:
	$(eval PG_URL?=$(DEFAULT_PG_URL))
	$(LOCAL_BIN)/goose -dir ./migrations postgres "$(PG_URL)" down

.PHONY: migration-status
migration-status:
	$(eval PG_URL?=$(DEFAULT_PG_URL))
	$(LOCAL_BIN)/goose -dir ./migrations postgres "$(PG_URL)" status

.PHONY: migration-create-sql
migration-create-sql:
	$(LOCAL_BIN)/goose -dir ./migrations create $(filter-out $@,$(MAKECMDGOALS)) sql

.PHONY: unit-tests
unit-tests:
	go test  ./internal/... ./pkg/...

.PHONY: run-all-tests
run-all-tests: unit-tests
	$(eval PG_URL?=$(DEFAULT_PG_URL))
	TEST_DATABASE_URL=$(PG_URL) go test ./tests/postgres/... -tags=integration

.PHONY: generate-swag
generate-swag:
	GOBIN=$(LOCAL_BIN) swag init ./internal/http/http.go

.PHONY: bin-deps
bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(LOCAL_BIN) go install github.com/vburenin/ifacemaker@latest
	GOBIN=$(LOCAL_BIN) go install go.uber.org/mock/mockgen@latest
	GOBIN=$(LOCAL_BIN) go install github.com/swaggo/swag/cmd/swag@latest
	GOBIN=$(LOCAL_BIN) go install gopkg.in/reform.v1/...
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest

.PHONY: generate
generate:
	rm -rfd internal/repo/schema/pgmodel/*
	cd internal/repo/schema/pgmodel && $(LOCAL_BIN)/reform-db -db-driver=postgres -db-source=$(DEFAULT_PG_URL) init
	cd internal/repo/schema/pgmodel && $(LOCAL_BIN)/reform ./
