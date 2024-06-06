# Use the official Golang image as the base image
FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Set environment variables
ENV MONGO_URL=mongodb://mongo:27017/godb

# Expose the application port
EXPOSE 6000

# Command to run the application
CMD ["./main"]
