# Use the official Go image.
FROM golang:1.20-alpine

# Set the working directory inside the container.
WORKDIR /app

# Copy the Go Modules manifests.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the source code.
COPY . .

# Build the Go app.
RUN go build -o /my-ecommerce-app ./cmd/api

# Expose port 8080 to the outside world.
EXPOSE 8080

# Command to run the executable.
CMD ["/my-ecommerce-app"]
