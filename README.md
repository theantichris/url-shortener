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
1. `go run main.go`

#### With Redis

1. Start database
    * `docker-compose -f stack.yml up`
1. Copy `sample.env` to `.env` and update values
1. `go run main.go`

## TODOs

* write tests
* deploy to Heroku
