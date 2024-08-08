# Go HTTP Server with Docker

This repository contains a simple Go HTTP server and a Dockerfile for containerization.

## Go Program

The Go program does the following:

1. Imports necessary packages: `fmt` for console output and `net/http` for HTTP server functionality.
2. In the `main` function:
   - Registers the `helloHandler` function to handle requests to the root path ("/").
   - Prints a message indicating the server is running.
   - Starts the server on port 8080.
   - Prints any error messages if the server fails to start.
3. The `helloHandler` function:
   - Takes two parameters: `w` (http.ResponseWriter) and `r` (*http.Request).
   - Writes "Hello, World!" to the response.

## Dockerfile

The Dockerfile uses a multi-stage build process:

### Stage 1: Builder

1. Uses `golang:1.20-alpine` as the base image.
2. Sets the working directory to `/app`.
3. Copies Go source files from `../src/` into the container.
4. Initializes a Go module named "main".
5. Builds the Go application with Linux-specific flags.

### Stage 2: Final Image

1. Starts from `alpine:latest`.
2. Installs CA certificates.
3. Sets the working directory to `/root/`.
4. Copies the compiled binary from the builder stage.
5. Exposes port 8080.
6. Specifies the command to run the executable.