# Use Golang image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Copy go files
COPY . .

# Download dependencies
RUN go mod tidy

# Build the Go app
RUN go build -o backend cmd/main.go

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./backend"]
