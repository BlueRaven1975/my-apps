# Use the official Go image from the Docker Hub to build the application
FROM golang:latest AS build

# Set the working directory inside the container
WORKDIR /app

# Copy dependencies management file and download/verify them
COPY go.mod .
RUN go mod tidy

# Copy the application directory contents into the container
COPY . .

# Build the application
RUN go build -o go-hello-world

# Use Alpine as a new base to shrink the final image size
FROM alpine:latest

# Add glibc compatibility libraries
RUN apk add libc6-compat

# Copy the application built in the previous stage
COPY --from=build /app/go-hello-world .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the application when the container launches
CMD ["./go-hello-world"]