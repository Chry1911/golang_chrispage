# Start from the official golang image for Go 1.22.2
FROM golang:1.22.2 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /godocker

EXPOSE 8080

CMD ["/godocker"]