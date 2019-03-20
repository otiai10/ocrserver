# ocrserver

[![Build Status](https://travis-ci.org/otiai10/ocrserver.svg?branch=master)](https://travis-ci.org/otiai10/ocrserver)
[![codecov](https://codecov.io/gh/otiai10/ocrserver/branch/master/graph/badge.svg)](https://codecov.io/gh/otiai10/ocrserver)
[![Go Report Card](https://goreportcard.com/badge/github.com/otiai10/ocrserver)](https://goreportcard.com/report/github.com/otiai10/ocrserver)

Simple OCR server, as a small working sample for [gosseract](https://github.com/otiai10/gosseract).

Try now here https://ocr-example.herokuapp.com/, and deploy your own now.

[![](https://user-images.githubusercontent.com/931554/36279290-7134626a-124b-11e8-8e47-d93b7122ea0d.png)](https://ocr-example.herokuapp.com)

# Deploy to Heroku

```sh
# Get the code
% git clone git@github.com:otiai10/ocrserver.git
% cd ocrserver
# Make your app
% heroku login
% heroku create
# Deploy the container
% heroku container:login
% heroku container:push web
# Enjoy it!
% heroku open
```

cf. [heroku cli](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)


# Quick Start

## Ready-Made Docker Image

```sh
% docker run -p 8080:8080 otiai10/ocrserver
# open http://localhost:8080
```

cf. [docker](https://www.docker.com/products/docker-toolbox)

## Development with Docker Image

```sh
% docker-compose up
# open http://localhost:8080
```

cf. [docker-compose](https://www.docker.com/products/docker-toolbox)

## Manual Setup

If you have tesseract-ocr  and library files on your machine

```sh
% go get github.com/otiai10/ocrserver/...
% PORT=8080 ocrserver
# open http://localhost:8080
```

cf. [gosseract](https://github.com/otiai10/gosseract)

# Documents

- [API Endpoints](https://github.com/otiai10/ocrserver/wiki/API-Endpoints)
