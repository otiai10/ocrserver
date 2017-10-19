
# ocrserver

Simple OCR server, as a small working sample for [gosseract](https://github.com/otiai10/gosseract).

Try now here https://ocr-example.herokuapp.com/, and deploy your own now.

# Quick Start

## Ready-Made Image

```sh
% docker run -it --rm \
  -e PORT=7777 -p 8080:7777 \
  otiai10/ocrserver
```

cf. [docker](https://www.docker.com/products/docker-toolbox)

## Local Development

```sh
% docker-compose up
# open http://localhost:8080
```

cf. [docker-compose](https://www.docker.com/products/docker-toolbox)

## Deploy to Heroku

```sh
% heroku create
% heroku container:login # If needed
% heroku container:push web
# heroku open
```

cf. [heroku cli](https://devcenter.heroku.com/articles/heroku-cli#download-and-install)

# Documents

- [API Endpoints](https://github.com/otiai10/ocrserver/wiki/API-Endpoints)
