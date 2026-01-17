FROM golang:1.25-alpine
# Install git (needed if go mod needs to fetch modules)
RUN apk add --no-cache git

# Set working directory
WORKDIR /app

# Copy Go modules first for caching
COPY go.mod ./


# Copy the rest of the code
COPY . .

# Expose API port
EXPOSE 8081

# Run the Go app
CMD ["go", "run", "cmd/server/main.go"]
