FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .
COPY ./.env.example ./.env

# Build the Go application
RUN go build -o app ./cmd/server/main.go

# Create a minimal production image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/app .

# Expose the application port
EXPOSE 8080

# Command to run the executable
CMD ["./app"]