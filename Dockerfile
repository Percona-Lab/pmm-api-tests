FROM golang:1.11.0-alpine3.7

RUN apk update && \
  apk upgrade --update-cache --available
RUN apk add git make curl perl bash build-base zlib-dev ucl-dev

RUN mkdir -p $GOPATH/src/github.com/Percona-Lab/pmm-api-tests

WORKDIR $GOPATH/src/github.com/Percona-Lab/pmm-api-tests/
COPY . $GOPATH/src/github.com/Percona-Lab/pmm-api-tests/

CMD make all && make run