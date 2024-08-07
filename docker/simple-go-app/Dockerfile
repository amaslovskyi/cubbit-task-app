# Start from the official golang image
FROM golang:1.20-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY ../src/*.go ./

# Initialize Go module
RUN go mod init main

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /main

# Start a new stage from scratch
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /main .

# Expose port 8080 to the outside world
EXPOSE 8081

# Command to run the executable
CMD ["./main"]
