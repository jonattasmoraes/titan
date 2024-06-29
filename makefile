.PHONY: default run

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
protoc:
	@protoc --go_out=. --go-grpc_out=. ./proto/user.proto        
clean:
	@rm -rf docs