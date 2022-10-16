# syntax=docker/dockerfile:1

##
## Build
##
FROM golang:1.19 AS build

# Define the working directory
WORKDIR /app

# Copy all files to container
COPY . .

# Install all dependencies
RUN go mod tidy

# Build the go project to binary
RUN go build -o /app/album-api main.go

##
## Deploy
##
FROM gcr.io/distroless/base-debian10

# Define the working directory
WORKDIR /app

# Copy the env
COPY --from=build /app/.env /app/.env

# Copy the binary that has been built
COPY --from=build /app/album-api /app/album-api

# Expose port 4000, then localhost can access the container
EXPOSE 4000

# Change user to nonroot
USER nonroot:nonroot

# Run the binary build of go project
ENTRYPOINT ["/app/album-api"]