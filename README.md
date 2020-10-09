# url-shortener

![Go](https://github.com/theantichris/url-shortener/workflows/Go/badge.svg) [![PkgGoDev](https://pkg.go.dev/badge/github.com/theantichris/url-shortener)](https://pkg.go.dev/github.com/theantichris/url-shortener)

A URL shortener written in Go.

## Available Ports & Adaptors

### Repositories

* MongoDB
* Redis

### Serializers

* JSON
* MessagePack

## Development

1. Start Mongo and Redis servers
    * `docker-compose up`
    * For Mongo add the `redirects` collection
    * Mongo Express is accessible at `http://localhost:8081`
1. Copy `sample.env` to `.env` and update values
    * Set `DB_TYPE` to `mongo` or `redis`
1. Start the application
    * `go run main.go`
    * The application is accessible at `http://localhost:8080`

## Endpoints

### GET /

A health check endpoint that displays a message if the server is running.

### POST /

Creates a new redirect and returns the code.

#### Body

```json
{
    "url": "http:/example.com"
}
```

#### Response

```json
{
  "code": "Rz5x645Mg",
  "url": "https://example.com",
  "created_at": 1602164689
}
```

### GET /{code}

Redirects you to the URL for the specified code.
