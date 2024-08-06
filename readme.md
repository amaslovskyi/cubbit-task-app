This Go program does the following:

1.We import the necessary packages: fmt for printing to the console and net/http for creating the HTTP server.
2.In the main function:
We use http.HandleFunc("/", helloHandler) to register the helloHandler function to handle requests to the root path ("/").
We print a message to the console indicating that the server is running.
We start the server using http.ListenAndServe(":8080", nil), which makes the server listen on port 8080.
If there's an error starting the server, we print the error message.
3.The helloHandler function is defined to handle incoming HTTP requests:
It takes two parameters: w (http.ResponseWriter) to write the response, and r (*http.Request) which contains information about the request.
It uses fmt.Fprintf(w, "Hello, World!") to write the "Hello, World!" message to the response.

Docker

This Dockerfile defines a multi-stage build process for a Go application:
1.The first stage uses the golang:1.20-alpine image as the base and is named "builder".
2.It sets the working directory to /app.
3.It copies Go source files from the ../src/ directory into the container.
4.It initializes a Go module named "main".
5.It builds the Go application with specific flags for a Linux environment.

The second stage starts from the alpine:latest image:
1.It installs CA certificates.
2.Sets the working directory to /root/.
3.Copies the compiled binary from the first stage.
4.Exposes port 8080.
5.Specifies the command to run the executable.

docker build -f docker/simple-go-app/Dockerfile -t shatten/cubbit-task-app:1.1 .
docker push shatten/cubbit-task-app:1.1

