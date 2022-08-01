# syntax=docker/dockerfile:1
FROM golang:1.18-alpine

ARG DEFAULT_PORT="8080"
WORKDIR /app

COPY go.mod ./
COPY *.go ./

RUN go build -o /go-public-ip
ENV PORT=${DEFAULT_PORT}
EXPOSE $PORT

CMD [ "/go-public-ip" ]