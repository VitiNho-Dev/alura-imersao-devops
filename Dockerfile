# Use the official Golang image as a base
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download and verify dependencies
RUN go mod download && go mod verify

# Copy the source code into the container
COPY . .

# Build the application
RUN go build -o main ./api/main.go

# Use a minimal Alpine Linux image for the final image
FROM alpine:latest AS runner

# Install ca-certificates for secure HTTPS communication
RUN apk --no-cache add ca-certificates

# Set the working directory in the final image
WORKDIR /app

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Expose the port the application listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
