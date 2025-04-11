# 📦 Mercado Livre - Desafio Técnico

Este repositório contém a solução para um desafio técnico com foco em alocação de pedidos em centros de distribuição. O sistema é composto por duas APIs independentes, que se comunicam entre si:

- **order-api** (Go - Gin-Gonic): Responsável por receber pedidos e alocar itens para os centros de distribuição apropriados.
- **distribution-center-api** (Python - FastAPI): Fornece informações aleatórias dos centros de distribuição e os itens disponíveis em cada um.

## 📚 Índice

- [📦 Descrição](#-mercado-livre---desafio-técnico)
- [🧠 Visão Geral](#-visão-geral)
- [🗂 Estrutura do Projeto](#-estrutura-do-projeto)
- [🚀 Como Executar Localmente](#-como-executar-localmente)
  - [Pré-requisitos](#pré-requisitos)
  - [Subir os serviços](#subir-os-serviços)
  - [Executar APIs manualmente](#alternativa-executar-apis-manualmente)
- [🔍 Exemplos de Uso](#-exemplos-de-uso)
- [🧪 Testes](#-testes)
- [📄 Documentação](#-documentação)
- [🧰 Tecnologias Utilizadas](#-tecnologias-utilizadas)
- [🧠 Lógica de agrupamento de Item/CD](#-lógica-de-agrupamento-de-itemcd)
- [📬 Contato](#-contato)

## 🧠 Visão Geral

Quando um pedido é recebido pela `order-api`, o sistema consulta a `distribution-center-api` para obter os centros que possuem os itens do pedido. Em seguida, aloca os itens da forma mais eficiente possível, retornando a lista de centros de distribuição utilizados.

---

## 🗂 Estrutura do Projeto

```
mercado-livre-desafio-tecnico/
│
├── order-api/                       # API de pedidos (Go)
│   ├── cmd/server/                  # Entrypoint
│   ├── config/                      # Configurações
│   ├── internal/                    # Domínio, handlers e rotas
│   ├── pkg/                         # Bibliotecas auxiliares (HTTP client, logger)
│   ├── docs/                        # Documentação Swagger
│   └── Makefile
│
├── distribution-center-api/         # API dos centros de distribuição (Python)
│   ├── app/
│   │   ├── controllers/             # Endpoints da API
│   │   └── services/                # Regras de negócio
│   └── pyproject.toml / requirements.txt
│
└── .docker/                         # Dockerfiles e docker-compose
```

---

## 🚀 Como Executar Localmente

### Pré-requisitos

- Docker + Docker Compose

### Subir os serviços

```bash
docker compose -f .docker/docker-compose.yml up --build -d
```

Isso irá iniciar:

- `distribution-center-api` na porta `8001`
- `order-api` na porta `8080`
- `mongodb` na porta `27017`

### Alternativa: Executar APIs manualmente

#### `order-api` (Go)

```bash
cd order-api
make run
```

#### `distribution-center-api` (Python)

```bash
cd distribution-center-api
poetry install
poetry run flask run
```

---

## 🔍 Exemplos de Uso

### 1. Enviar pedido para processamento

`POST /api/v1/orders`

```json
{
  "items": [
    {
      "id": 1,
      "name": "Produto X",
      "price": 1.99
    }
  ]
}
```

**Resposta:**

```json
{
  "order_id": "67f8717595ea8b62288e5e88", // ID gerado automaticamente
  "order": {
    "items": [
      {
        "id": 1,
        "name": "Produto X",
        "price": 1.99,
        "distribution_center": "CD1"
      }
    ]
  }
}
```

### 2. Listar itens por centro

`GET /api/v1/orders/:id:`

**Resposta**

```json
{
  "order_id": "67f8717595ea8b62288e5e88", // ID gerado automaticamente
  "order": {
    "items": [
      {
        "id": 1,
        "name": "Produto X",
        "price": 1.99,
        "distribution_center": "CD1"
      }
    ]
  }
}
```

### 3. Listar itens por centro (na API Python)

`GET /distribuitioncenters?itemId=:id:`

**Resposta**

```json
{
  "distribuitionCenters": ["CD1", "CD2", "CD3"]
}
```

---

## 🧪 Testes

Toda a lógica de negócio e os algoritmos principais implementados neste projeto foram devidamente cobertos por testes unitários. Os testes garantem o correto funcionamento da lógica de alocação de pedidos, validação dos dados e comunicação entre os serviços, proporcionando maior segurança, confiabilidade e facilidade de manutenção do código.

### order-api

```bash
cd order-api
make test
```

---

## 📄 Documentação

### Swagger

- `order-api`: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

### Collection Insomnia v5

- [Collection Insomnia v5](./collection.yaml)

## 🧰 Tecnologias Utilizadas

- **Go + GinGonic** (API de pedidos)
- **Python + Flask** (API de centros de distribuição)
- **Docker + Docker Compose**
- **MongoDB** (Banco de dados não relacional)
- **Swagger/OpenAPI (via swaggo)**
- **Poetry** para gerenciamento de dependências Python

---

## 🧠 Lógica de agrupamento de Item/CD

Foi construído um algoritmo de complexidade O(n \* m) — linear em relação ao número de itens do pedido (n) e ao número de centros de distribuição (m) — para resolver o problema de Set Cover. Esse algoritmo percorre todos os itens do pedido e avalia, para cada item, os centros de distribuição disponíveis, garantindo uma alocação eficiente e direta, adequada para cenários em que simplicidade e performance prática são prioridades.

## 📬 Contato

email: marcioedumartinez@gmail.com
celular: (11) 94256-2000
