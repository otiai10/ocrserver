FROM golang:1.11

LABEL maintainer="otiai10 <otiai10@gmail.com>"

RUN apt-get -qq update
RUN apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr

# Load languages
RUN apt-get install -y \
  tesseract-ocr-jpn

ADD . $GOPATH/src/github.com/otiai10/ocrserver
WORKDIR $GOPATH/src/github.com/otiai10/ocrserver
RUN go get ./...
RUN go get -t github.com/otiai10/gosseract
RUN go test -v github.com/otiai10/gosseract

CMD $GOPATH/bin/ocrserver
