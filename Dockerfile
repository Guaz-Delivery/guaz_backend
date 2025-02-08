# using golang:1.24rc2bookworm as a base image
FROM golang:1.23-bookworm AS  base

# Move current working directory to /build
WORKDIR /build

# Copy go.mod and go.sum file to working directory
COPY go.mod go.sum ./

# Install Dependences
RUN go mod download

# Copy source code to the container
COPY . .

# Build the application
RUN go build -o guaz-hooks

# Exposing port 
EXPOSE 9999

# Starting the application 
CMD ["./guaz-hooks"]
