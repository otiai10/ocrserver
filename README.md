
# ocrserver

Simple OCR server, as a small working sample for [gosseract](https://github.com/otiai10/gosseract).

Try now here https://ocr-example.herokuapp.com/, and deploy your own now.

# Deploy to Heroku

```sh
% git clone git@github.com:otiai10/ocrserver.git
% cd ocrserver
# heroku login (if needed)
% heroku create
# heroku container:login (If needed)
% heroku container:push web
# heroku open
```

cf. [heroku cli](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)


# Quick Start

## Ready-Made Docker Image

```sh
% docker run -e PORT=8080 -p 8080:8080 otiai10/ocrserver
# open http://localhost:8080
```

cf. [docker](https://www.docker.com/products/docker-toolbox)

## Development with Docker Image

```sh
% docker-compose up --build
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
