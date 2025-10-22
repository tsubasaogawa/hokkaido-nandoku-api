# --- Builder Stage ---
FROM golang:1.22-alpine as builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is important for creating a static binary
# -o bootstrap specifies the output file name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o bootstrap ./cmd/api

# --- Final Stage ---
FROM alpine:latest

WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /app/bootstrap .

# Copy the data file
COPY --from=builder /app/data/nandoku_chimei.csv ./data/nandoku_chimei.csv

# This command is not strictly necessary for Lambda, 
# but it's good practice to define the entrypoint.
CMD ["./bootstrap"]
