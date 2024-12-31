# Step 1: Build the Go binary
FROM golang:1.23.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the entire project to the container's workspace
COPY . .

# Build the Go app
RUN go build -o main ./cmd

# Step 2: Run the Go app directly from the Go image
FROM golang:1.23.2

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary file from the builder stage
COPY --from=builder /app/main .

# Expose port that the Go app listens on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
