FROM golang:1.8

MAINTAINER otiai10 <otiai10@gmail.com>

RUN apt-get -qq update
RUN apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr

ADD . $GOPATH/src/github.com/otiai10/ocrserver
WORKDIR $GOPATH/src/github.com/otiai10/ocrserver
RUN go get ./...

CMD $GOPATH/bin/ocrserver
