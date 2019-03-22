FROM golang:1.12-alpine

RUN apk update && \
  apk upgrade --update-cache --available
RUN apk add git make build-base

RUN mkdir -p $GOPATH/src/github.com/Percona-Lab/pmm-api-tests

WORKDIR $GOPATH/src/github.com/Percona-Lab/pmm-api-tests/
COPY . $GOPATH/src/github.com/Percona-Lab/pmm-api-tests/

CMD make run
