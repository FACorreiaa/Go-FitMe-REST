# Start from golang base image
FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="a11199"

# Install git
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy the source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the dev.env file
#COPY dev.env .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
