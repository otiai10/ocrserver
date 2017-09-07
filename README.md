
# ocrserver

Simple OCR server, as a small working sample for [gosseract](https://github.com/otiai10/gosseract).

Try now here https://ocr-example.herokuapp.com/, and deploy your own now.

# Quick Start

## Local Development

with [docker](https://www.docker.com/products/docker-toolbox) and [docker-compose](https://www.docker.com/products/docker-toolbox) required

```sh
% docker-compose up
# open http://localhost:8080
```

## Deploy to Heroku

with also [heroku cli](https://devcenter.heroku.com/articles/heroku-cli#download-and-install) required

```sh
% heroku create
% heroku container:login # If needed
% heroku container:push web
# heroku open
```

# Documents

- [API Endpoints](https://github.com/otiai10/ocrserver/wiki/API-Endpoints)
