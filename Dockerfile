# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download the dependencies
COPY go.mod go.sum ./
RUN apk add gcc g++
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -tags musl -a -o server .

# Expose the ports for the server and MySQL
EXPOSE 8080

# Start the server and database
CMD ["sh", "-c", "/app/server"]