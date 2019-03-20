FROM golang:1.12
ENV GO111MODULE=on

LABEL maintainer="otiai10 <otiai10@gmail.com>"

RUN apt-get -qq update \
  && apt-get install -y \
    libleptonica-dev \
    libtesseract-dev \
    tesseract-ocr

# Load languages
RUN apt-get install -y \
  tesseract-ocr-jpn

ADD . $GOPATH/src/github.com/otiai10/ocrserver
WORKDIR $GOPATH/src/github.com/otiai10/ocrserver
RUN go get ./...
RUN go test -v github.com/otiai10/gosseract

ENV PORT=8080
CMD $GOPATH/bin/ocrserver
