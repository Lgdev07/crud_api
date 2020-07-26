FROM golang:alpine

WORKDIR /crud_api

COPY go.mod go.sum ./

RUN go mod download

COPY . .
