ifneq (,$(wildcard ./.env))
    include .env
    export
endif


migrate:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URL) goose -dir=$(MIGRATION_PATH) up

rollback:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URL) goose -dir=$(MIGRATION_PATH) down

new_migration:
	@GOOSE_DRIVER=postgres GOOSE_DBSTRING=$(DB_URL) goose -dir=$(MIGRATION_PATH) create $(name) sql
