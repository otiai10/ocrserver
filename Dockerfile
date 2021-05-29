FROM debian:bullseye-slim
LABEL maintainer="otiai10 <otiai10@gmail.com>"

ARG LOAD_LANG=jpn

RUN apt update \
    && apt install -y \
      ca-certificates \
      libtesseract-dev=4.1.1-2.1 \
      tesseract-ocr=4.1.1-2.1 \
      golang=2:1.15~1

ENV GO111MODULE=on
ENV GOPATH=${HOME}/go
ENV PATH=${PATH}:${GOPATH}/bin

ADD . $GOPATH/src/github.com/otiai10/ocrserver
WORKDIR $GOPATH/src/github.com/otiai10/ocrserver
RUN go get -v ./... && go install .

# Load languages
RUN if [ -n "${LOAD_LANG}" ]; then apt-get install -y tesseract-ocr-${LOAD_LANG}; fi

ENV PORT=8080
CMD ["ocrserver"]
