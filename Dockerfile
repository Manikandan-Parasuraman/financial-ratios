# Use a lightweight Go image
FROM golang:1.20-alpine

# Set up the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
