# Desafio Clean Architecture

## Pré-requisitos

- Docker
- Go

## Comando Necessários para Executar o Sistema

1. Subir os containers do docker:

```bash
docker compose up -d

```

1. Rodar o Makefile:

```bash
make migrate
```

1. No diretório cmd/ordersystem rodar o arquivo go  dos servidores com wire:

```bash
go run main.go wire_gen.go
```

## Portas Utilizadas

- Servidor Web na porta 8000
- Servidor gRPC na porta 50051
- Servidor GraphQL na porta 8080

## Modo de acesso de Sistema

- Requisições para o servidor Web usar cUrl ou os arquivos http dentro da pasta Api
- Requisições para o servidor gRPC utilizar Evans
- Requisições para o servidor GraphQL  utilizar o playground do graphql, exemplos no arquivo queries.graphql na pasta api.
