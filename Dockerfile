FROM golang:1.24-alpine

RUN apk add --no-cache curl git bash make

ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin
ENV GOFLAGS=-buildvcs=false

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
    sh -s -- -b $(go env GOPATH)/bin latest

WORKDIR /project

RUN mkdir -p /project/code

WORKDIR /project

COPY . .
