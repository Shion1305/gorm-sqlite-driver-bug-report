.PHONY: setup migrate_up create_migration

setup:
	go install -tags 'sqlite3' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

.PHONY: migrate_up
migrate_up:
	migrate -path migrations -database sqlite3://db.sqlite3 up

create_migration:
ifndef NAME
	@echo "Usage: make NAME=migration_name create_migration"
else
	migrate create -ext sql -dir db/migrations -seq $(NAME)
endif
