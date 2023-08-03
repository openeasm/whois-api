# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the necessary files into the container
COPY app.go ./

# Build the Go application inside the container
RUN go build -o app app.go

# Expose the port on which the application will listen
EXPOSE 8000

# Command to run the application when the container starts
CMD ["./app", "8000"]