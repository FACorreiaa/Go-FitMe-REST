## Start from golang base image
#FROM golang:alpine as builder
#
## Add Maintainer info
#LABEL maintainer="a11199"
#
## Install git
#RUN apk update && apk add --no-cache git
#
## Set the current working directory inside the container
#WORKDIR /app
#
## Download dependencies
#COPY go.mod go.sum ./
#RUN go mod download
#
## Copy the source code
#COPY . .
#
## Build the Go app
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
#
## Start a new stage from scratch
#FROM alpine:latest
#RUN apk --no-cache add ca-certificates
#
#WORKDIR /
#
## Copy the Pre-built binary file from the previous stage
#COPY --from=builder /app/main .
#
## Copy the .env file
#
## Expose port 8080 to the outside world
#EXPOSE 8080
#
## Command to run the executable
#CMD ["./main"]

FROM golang:1.20.6 as base

## Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM base as dev

## Install the air binary so we get live code-reloading when we save files
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Run the air command in the directory where our code will live
WORKDIR /opt/app/api
CMD ["air"]
