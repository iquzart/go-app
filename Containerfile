# Start from the latest golang base image
FROM golang:alpine

# Add Maintainer Info
LABEL maintainer="Muhammed Iqbal <iquzart@hotmail.com>"

ENV GIN_MODE=release

ENV BANNER="Hello from Go App"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o app .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"]