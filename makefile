.PHONY: default run build test docs clean

APP_NAME = titan

default: run

run:
	@docker-compose up -d
build:
	go build main.go
test:
	go test ./...
docs:
	@swag init --output docs --dir ./cmd/user,./internal/user/infra/http,./internal/user/domain/dto
clean:
	@rm -rf docs