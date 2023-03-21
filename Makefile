LOCAL_BIN:=$(CURDIR)/bin

install-minimock:
	GOBIN=$(LOCAL_BIN) go install github.com/gojuno/minimock/v3/cmd/minimock

run:
	sudo docker compose up --force-recreate --build

migrate:
	./migration.sh