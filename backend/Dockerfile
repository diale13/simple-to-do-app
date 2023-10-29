# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application code into the container
COPY . .

# Install any necessary Go dependencies (if you have any)
RUN go mod tidy

# Build the Go application
RUN go build -o app

EXPOSE 9000

# Command to run the application
CMD ["./app"]
