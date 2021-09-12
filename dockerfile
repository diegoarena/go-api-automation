FROM golang:1.15-alpine 

ENV CGO_ENABLED=0

WORKDIR /test

RUN set -ex; \
    apk update; \
    apk add --no-cache git 

RUN go get github.com/githubnemo/CompileDaemon
RUN go get gotest.tools/gotestsum

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

