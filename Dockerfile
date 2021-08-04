FROM debian:bullseye-slim AS build
LABEL maintainer="otiai10 <otiai10@gmail.com>"

RUN apt update \
    && apt install -y \
      ca-certificates \
      libtesseract-dev=4.1.1-2.1 \
      tesseract-ocr=4.1.1-2.1 \
      golang=2:1.15~1

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
ENV GOPATH=${HOME}/go
ENV PATH=${PATH}:${GOPATH}/bin

ADD . $GOPATH/src/github.com/otiai10/ocrserver
WORKDIR $GOPATH/src/github.com/otiai10/ocrserver
RUN go get -v ./... && go install .

# Load languages
ARG LOAD_LANG=chi-sim
RUN if [ -n "${LOAD_LANG}" ]; then apt-get install -y tesseract-ocr-${LOAD_LANG}; fi

ENV PORT=8080

FROM scratch AS OS

COPY --from=build /go/bin/ocrserver /
CMD [ "/ocrserver" ]