FROM golang:latest

LABEL version="1.0"

RUN mkdir /go/src/cartman
COPY . /go/src/cartman
WORKDIR /go/src/cartman

RUN go get

CMD go run main.go
