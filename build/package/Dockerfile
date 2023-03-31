# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

# stage BUILD
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o estimation-ceremony-backend main.go

# stage SETUP
ENV HOST=0.0.0.0
ENV PORT=8000
EXPOSE 8000
WORKDIR /app
ENTRYPOINT ./estimation-ceremony-backend run -h ${HOST} -p ${PORT}
