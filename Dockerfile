#Docker file
#FROM golang:1.19-alpine3.16 AS builder
#WORKDIR /app
#COPY . .
#RUN go build -o main main.go
#
## Run stage
#FROM alpine:3.16
#WORKDIR /app
#COPY --from=builder /app/main .
#
#EXPOSE 8080
#CMD [ "/app/main" ]

# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="a11199"

# Install git.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
