# Stage 1: Сборка (использует golang образ)
FROM golang:1.26.1-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o app cmd/main.go

# Stage 2: Запуск (использует легковесный alpine)
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]