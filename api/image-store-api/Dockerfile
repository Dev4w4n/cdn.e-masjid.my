# Stage 1: Build the application
FROM golang:alpine3.19 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container
COPY . .
COPY ./images/main.jpg .

# Install any dependencies your application might need
RUN go get -d -v ./...
RUN go install -v ./...

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o image-server

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Define build-time argument for GO_ENV
ARG GO_ENV=local

# Set Go environment variables
ENV GO_ENV=$GO_ENV

# Set the working directory inside the container
WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/image-server .
COPY --from=builder /app/.env.* .
COPY --from=builder /app/main.jpg .

RUN mkdir -p /app2

# Expose the port that your application will run on
EXPOSE 8080

# Command to run your application
CMD ["./image-server", "-env", "$GO_ENV"]
