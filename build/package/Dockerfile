# syntax=docker/dockerfile:1

FROM golang:1.20

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o estimation-ceremony-backend ./cmd/estimation-ceremony-backend/main.go

EXPOSE 8000

ENTRYPOINT ["./estimation-ceremony-backend", "--domain", "localhost", "--port", "8000"]
