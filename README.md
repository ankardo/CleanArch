# Desafio Clean Architecture

## Pré-requisitos

- Docker
- Go

## Comandos Necessários para Executar o Sistema

### Execução Inicial

```bash
docker compose up --build -d
```

### Execuções Futuras

```bash
docker compose up -d
```

## Portas Utilizadas

- Servidor Web na porta 8000
- Servidor gRPC na porta 50051
- Servidor GraphQL na porta 8080

## Modo de acesso de Sistema

- Requisições para o servidor Web usar cUrl ou os arquivos http dentro da pasta Api
- Requisições para o servidor gRPC utilizar Evans
- Requisições para o servidor GraphQL  utilizar o playground do graphql, exemplos no arquivo queries.graphql na pasta api.
