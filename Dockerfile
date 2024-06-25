FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .env ./

RUN go install github.com/air-verse/air@latest

COPY . .

WORKDIR /app