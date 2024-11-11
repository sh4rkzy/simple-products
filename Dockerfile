FROM golang:1.18-alpine AS build

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o rest-go .

FROM alpine:latest AS production-stage

WORKDIR /root/

COPY --from=build /app/rest-go .

EXPOSE 8080

CMD ["./myapp"]
