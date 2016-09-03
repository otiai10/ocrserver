FROM golang:1.7

MAINTAINER otiai10 <otiai10@gmail.com>

RUN apt-get -qq update
RUN apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr

ADD . /go/src/github.com/otiai10/ocrserver
WORKDIR /go/src/github.com/otiai10/ocrserver
RUN go get ./...

ENTRYPOINT /go/bin/ocrserver
