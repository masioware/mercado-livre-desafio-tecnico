# Diretório de saída do binário
BINARY_NAME=order-api
MAIN=./cmd/server

.PHONY: all build run test lint tidy clean

all: test build

## Builda o binário principal
build:
	go build -o bin/$(BINARY_NAME) $(MAIN)

## Roda a aplicação localmente
run:
	go run $(MAIN)

## Executa todos os testes com verbose
test:
	go test ./... -v

## Executa o linter (requer golangci-lint instalado)
lint:
	golangci-lint run

## Garante que o go.mod/go.sum estão atualizados
tidy:
	go mod tidy

## Limpa binário
clean:
	rm -rf bin/

## Cria arquivos do swagger
swagger:
	swag init --parseDependency -g cmd/server/main.go