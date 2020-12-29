# Start from the latest golang base image
FROM golang:1.15-alpine as  build-env

# Set the Current Working Directory inside the container
WORKDIR /src

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o app .

# Final stage
FROM alpine:latest

# Updated the Image
RUN apk update && rm -rf /var/cache/apk/*

# Maintainer Info
LABEL maintainer="Muhammed Iqbal <iquzart@hotmail.com>"

# Set GIN Mode as Release
ENV GIN_MODE=release

# Set environment variable default value
ENV BANNER="Hello from Go App"

# Container PORT
ENV PORT="8080"

# Create non-root account to run the container
RUN adduser go -h /app -u 1000 -D

# Switch to non-root user
USER go

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy app
COPY --from=build-env /src/app .

# Copy public HTML files
COPY public public

# Expose port 8080 to the outside world
EXPOSE ${PORT}

#ENTRYPOINT ["./app"]
CMD ./app
