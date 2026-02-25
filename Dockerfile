FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY config/config.yaml /app/config/
CMD ["./server"]