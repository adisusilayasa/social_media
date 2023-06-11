# Start from a base Golang image
FROM golang:alpine AS builder

RUN apk update &&\
    apk add --no-cache ca-certificates tzdata\
    && update-ca-certificates

# Install MySQL client
RUN apk add --no-cache mysql-client

# Set the working directory
WORKDIR /social_media

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the workspace
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o social_media ./cmd/app

# Create a new stage with a minimal image
FROM alpine:latest

RUN apk --no-cache add tzdata

# Set the working directory to /app
WORKDIR /social_media

# Copy the binary from the builder stage
COPY --from=builder /social_media/social_media /social_media
COPY config/config.yaml.example /social_media/config/config.yaml


# Expose port 8080
EXPOSE 80
EXPOSE 443

# Run the binary
ENTRYPOINT ["./social_media"]
