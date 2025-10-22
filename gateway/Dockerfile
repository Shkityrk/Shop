# Build stage
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install swag CLI for generating swagger docs
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Copy everything
COPY . .

# Download dependencies and create go.sum
RUN go mod download && go mod tidy

# Generate swagger documentation
RUN swag init

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o gateway .

# Final stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the binary from builder
COPY --from=builder /app/gateway .

# Expose port
EXPOSE 8000

# Run the application
CMD ["./gateway"]

