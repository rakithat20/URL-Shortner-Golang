FROM golang:1.23.2

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy