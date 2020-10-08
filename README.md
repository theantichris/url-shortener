# url-shortener

![Go](https://github.com/theantichris/url-shortener/workflows/Go/badge.svg)

A URL shortener written in Go.

## Ports

### Repositories

* MongoDB
* Redis

### Serializers

* JSON
* MessagePack

### Dev

#### With MongoDB

1. Start database
    * `docker-compose -f stack.yml up`
    * Go to `http://localhost:8081` for Mongo Express
    * Create a `redirects` collection
1. Copy `sample.env` to `.env` and update values

#### With Redis

1. Start database
    * `docker-compose -f stack.yml up`
1. Copy `sample.env` to `.env` and update values

## TODOs

* write tests
* write documentation
* deploy to Heroku
