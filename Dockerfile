# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy the necessary Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/booking_platform cmd/web/main.go

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/booking_platform .

# Expose the port your application runs on
EXPOSE 9998

# Run the application
CMD ["./booking_platform"]
