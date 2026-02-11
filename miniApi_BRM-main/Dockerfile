
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api


FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /app/main .
COPY ca.pem /app/internal/db/certs/ca.pem
COPY .env .env


ENV SERVER_PORT=8080
ENV DB_TLS_MODE=custom
ENV DB_CA_PATH=/app/internal/db/certs/ca.pem

EXPOSE 8080
CMD ["./main"]
