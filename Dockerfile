# Stage 1: Build the Go application
FROM golang:latest AS builder

WORKDIR /app

# Copy only the necessary files for building
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# # Build the Go application
# RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Stage 2: Create a lightweight runtime image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the compiled binary from the build stage
COPY --from=build /app/app .

# Expose the port your application runs on
EXPOSE 9998

# Run the application
CMD ["./app"]
