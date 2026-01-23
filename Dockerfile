# Use official Go image
FROM golang:1.18.3

# Set working directory
WORKDIR /usr/src/app

# Copy all project files
COPY . .

# Download dependencies
RUN go mod tidy

# Expose the port your app uses
EXPOSE 3000

# Run the application
CMD ["go", "run", "main.go"]
