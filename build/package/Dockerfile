# syntax=docker/dockerfile:1

FROM golang:1.20-alpine AS builder

WORKDIR /build
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o estimation-ceremony-backend main.go

FROM scratch

WORKDIR /app
COPY --from=builder /build/estimation-ceremony-backend /app/estimation-ceremony-backend
EXPOSE 8000
ENTRYPOINT ["/app/estimation-ceremony-backend", "run", "-h", "0.0.0.0", "-p", "8000"]
