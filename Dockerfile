# Use the official Golang image to build the binary
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./
# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bootstrap ./cmd/api

# Use a minimal image to run the binary
FROM alpine:latest

WORKDIR /root/

# Copy the binary from the builder stage
COPY --from=builder /app/bootstrap .
# Copy the data file
COPY --from=builder /app/data/nandoku_chimei.csv ./data/nandoku_chimei.csv

# Expose port 8080
EXPOSE 8080

# Command to run the executable
CMD ["./bootstrap"]
