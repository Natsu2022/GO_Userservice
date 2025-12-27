# -------- Build stage --------
FROM golang:1.23-alpine AS builder

WORKDIR /app

# copy go mod first for caching
COPY go.mod go.sum ./
RUN go mod tidy

# copy source code
COPY . .

# build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o user_service ./cmd/user_service

# -------- Runtime stage --------
FROM alpine:3.19

WORKDIR /app

# copy binary from builder
COPY --from=builder /app/user_service .

EXPOSE 3455

CMD ["./user_service"]
