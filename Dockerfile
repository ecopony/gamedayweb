FROM golang:latest

ADD . /go/src/github.com/ecopony/gamedayweb

WORKDIR /go/src/github.com/ecopony/gamedayweb
RUN go get github.com/ecopony/gamedayapi
RUN go install
ENTRYPOINT /go/bin/gamedayweb

EXPOSE 3000
