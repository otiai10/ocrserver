# make base image.
FROM debian:bullseye-slim AS base
LABEL maintainer="otiai10 <otiai10@gmail.com>" \
	updatedBy="RocSun <oldsixa@163.com>"

RUN apt-get update \
    && apt-get install -y ca-certificates \
      apt-transport-https \
    # this setting local apt mirrors
    && cp -a /etc/apt/sources.list /etc/apt/sources.list.bak \ 
    && sed -i "s@http://deb.debian.org@https://repo.huaweicloud.com@g" /etc/apt/sources.list \
    && sed -i "s@http://security.debian.org@https://repo.huaweicloud.com@g" /etc/apt/sources.list \
    && apt-get update \
    && apt-get install -y libtesseract-dev=4.1.1-2.1 \
      tesseract-ocr=4.1.1-2.1

# build source.
FROM base AS build

RUN apt install -y golang=2:1.15~1

ENV GO111MODULE=on
ENV GOPROXY="https://goproxy.cn,direct"
ENV GOPATH=/go
ENV PATH=${PATH}:${GOPATH}/bin

ADD . $GOPATH/src/ocrserver
WORKDIR $GOPATH/src/ocrserver
RUN go get -v ./... && go install .

# make run env. This is build image.
FROM base AS OS
ENV PORT=8080
COPY --from=build /go /go

ARG LOAD_LANG=chi-sim
RUN if [ -n "${LOAD_LANG}" ]; then apt-get install -y tesseract-ocr-${LOAD_LANG}; fi \
    && apt-get clean
CMD [ "/go/bin/ocrserver" ]

# ################################################################################
# # debian:bullseye-slim image size of is about 220M.                            #
# # if use alpine. image size of is about 160M. Here's the alpine-based code     #
# # min image in dockerhub rocsun/ocr-server                                     #
# ################################################################################

# FROM alpine:3.14 as base
#LABEL maintainer="otiai10 <otiai10@gmail.com>" \
# 	 updatedBy="RocSun <oldsixa@163.com>"

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories \
#     && apk add tesseract-ocr \
#     && apk add tesseract-ocr-dev \
#     && apk add tesseract-ocr-data-chi_sim

# FROM base AS build

# RUN apk add go \
#       g++

# ENV GO111MODULE=on
# ENV GOPROXY="https://goproxy.cn,direct"
# ENV GOPATH=/go
# ENV PATH=${PATH}:${GOPATH}/bin

# ADD . $GOPATH/src/ocrserver
# WORKDIR $GOPATH/src/ocrserver
# RUN go get -v ./... && go install .

# FROM base AS OS
# RUN apk add bash
# COPY --from=build /go /go
# ENV PORT=8080
# CMD ["/go/bin/ocrserver"]
