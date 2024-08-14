FROM golang:1.22.6-alpine

WORKDIR /app

RUN apk add --no-cache bash

COPY go.mod go.sum ./

COPY . .

RUN go mod tidy && \
    go build -o main ./cmd/go-ama-queue && \
    chmod +x main
    
COPY scripts/wait-for-it.sh /usr/local/bin/wait-for-it.sh
RUN chmod +x /usr/local/bin/wait-for-it.sh

EXPOSE 8080
