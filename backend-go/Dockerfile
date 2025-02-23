# Use Golang image for building the binary
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy dependency files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the code
COPY . .

# Build the Go application (optimized)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o backend ./cmd/main.go

# Use a minimal runtime image
FROM gcr.io/distroless/static-debian12

# Copy the compiled binary
COPY --from=builder /app/backend /bin/backend

# Expose port
EXPOSE 8080

# Run the binary
CMD ["/bin/backend"]
