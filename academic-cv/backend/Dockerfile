# ---------- Build Stage ----------
FROM golang:1.25-alpine AS builder

# 1. Install build dependencies (needed for CGO/SQLite)
RUN apk add --no-cache git gcc musl-dev

# 2. Set working directory
WORKDIR /app

# 3. Copy go.mod and go.sum first
COPY go.mod go.sum ./

# 4. Set Go environment variables
ENV GOPROXY=https://proxy.golang.org,direct
ENV GOSUMDB=off
ENV CGO_ENABLED=1
ENV GOOS=linux

# 5. Download dependencies
RUN go mod tidy
RUN go mod download -x

# 6. Copy the rest of the source code
COPY . .

# 7. Build the binary
RUN go build -o server ./cmd/server/main.go

# ---------- Final Stage ----------
FROM alpine:latest

# 8. Install runtime dependencies (needed for SQLite)
RUN apk add --no-cache libgcc libstdc++ ca-certificates

WORKDIR /app

# 9. Copy the compiled binary from builder
COPY --from=builder /app/server .

# 10. Expose port
EXPOSE 8081

# 11. Run the server
CMD ["./server"]
