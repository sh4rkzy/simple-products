# Etapa de build
FROM golang:1.23-alpine AS build

# Instalar dependências necessárias
RUN apk add --no-cache git

WORKDIR /app

# Copiar arquivos de dependências primeiro
COPY go.mod go.sum ./
RUN go mod download

# Copiar todo o código-fonte para o container
COPY . .

# Compilar o binário
RUN go build -o app /app/cmd/main.go

# Etapa de produção
FROM alpine:latest AS production

WORKDIR /root/

# Copiar o binário gerado
COPY --from=build /app/app .

COPY .env /root/.env
# Expor a porta usada pela aplicação
EXPOSE 3000

# Comando de inicialização
CMD ["./app"]
