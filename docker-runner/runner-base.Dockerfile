FROM golang:1.9
# Base image that already has the pre-requisites downloaded.

WORKDIR /go/src/github.com/zmap

RUN go-wrapper download github.com/dmin12/zgrab_tls1.3

WORKDIR /go/src/github.com/dmin12/zgrab_tls1.3

RUN go get -v ./...
RUN go get -v -t ./...
