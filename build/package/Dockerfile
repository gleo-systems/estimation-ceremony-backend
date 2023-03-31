# syntax=docker/dockerfile:1

FROM golang:1.20

ENV HOST=0.0.0.0
ENV PORT=8000

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./
RUN go build -o estimation-ceremony-backend main.go

EXPOSE 8000

ENTRYPOINT ./estimation-ceremony-backend run -h ${HOST} -p ${PORT}
