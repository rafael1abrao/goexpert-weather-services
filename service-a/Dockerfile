# Etapa 1: build da aplicação
FROM golang:1.23-alpine AS builder

# Cria diretório de trabalho
WORKDIR /app

# Copia go.mod e go.sum primeiro (melhor uso de cache)
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante da aplicação
COPY . .

# Compila o binário estaticamente
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./main.go

# Etapa 2: imagem final minimalista
FROM alpine:latest

WORKDIR /root/

# Copia o binário da etapa anterior
COPY --from=builder /app/server .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando de execução
CMD ["./server"]