FROM golang:1.15-alpine 

ENV CGO_ENABLED=0

WORKDIR /test

RUN set -ex; \
    apk update; \
    apk add --no-cache git 

RUN go get github.com/githubnemo/CompileDaemon
RUN go get gotest.tools/gotestsum
RUN go get -u github.com/vakenbolt/go-test-report/

#Download project dependencies 
COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

