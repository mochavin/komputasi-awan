FROM golang:1.23-alpine AS builder

# Ensure tools for fetching modules are available
RUN apk add --no-cache git ca-certificates

WORKDIR /app

# Copy the full project first so 'go mod tidy' can generate go.sum and download deps
COPY . .

# Build environment
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# Ensure module files are tidy and dependencies are downloaded, then build
RUN go mod tidy
RUN go build -o /app/main .

# Small runtime image
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080
CMD ["./main"]
