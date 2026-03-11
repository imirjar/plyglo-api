# Stage 1: Builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first to leverage Docker cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 disables CGO, creating a statically linked binary
# -o specifies the output file name
RUN go build -o poliglotim-api cmd/main.go

EXPOSE 8080

# Command to run the application
CMD ["./poliglotim-api"]