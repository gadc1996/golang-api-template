# syntax=docker/dockerfile:1

FROM golang:1.20rc1-alpine3.17

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /api-template

EXPOSE 8080

CMD [ "/api-template" ]