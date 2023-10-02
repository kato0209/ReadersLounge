DOCKER_COMPOSE := docker compose


POSTGRES_USER := root
DATABASE      := readers_lounge_db
psql:
	$(DOCKER_COMPOSE) exec db psql -U $(POSTGRES_USER) -d $(DATABASE)

GOOSE_DRIVER   := postgres
GOOSE_DBSTRING ?= host=db user=root dbname=readers_lounge_db password=p@ssword sslmode=disable

migrate/status:
	$(DOCKER_COMPOSE) run --rm migration status

VERSION:=$(shell ls db/migration | awk -F"_*.sql" 'BEGIN {max=0} {split($$1, a, "_"); if(a[1]>max){max = a[1]}}END{print max+1}')
TEMPLATE?=
migrate/new:
	echo '-- +goose Up' > db/migration/${VERSION}_${TEMPLATE}.sql

migrate/up:
	$(DOCKER_COMPOSE) run --rm migration up