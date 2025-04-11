# Etapa de build
FROM golang:alpine AS builder

# Instala dependências essenciais
RUN apk add --no-cache git

# Cria diretório da aplicação
WORKDIR /usr/order-api

# Copia os arquivos da aplicação
COPY ./order-api/go.mod ./order-api/go.sum ./
RUN go mod download

COPY ./order-api/cmd ./cmd
COPY ./order-api/config ./config
COPY ./order-api/docs ./docs
COPY ./order-api/internal ./internal
COPY ./order-api/pkg ./pkg

# Compila o binário (static build)
RUN go build -o main ./cmd/server

# Etapa final (imagem enxuta)
FROM alpine:latest

WORKDIR /root/

# Copia o binário da etapa de build
COPY --from=builder /usr/order-api/main .

# Expõe a porta
EXPOSE 8080

# Comando de entrada
CMD ["./main"]