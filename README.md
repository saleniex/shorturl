Short URL
==

URL shortening webservice with REST API for URL management. 

## Set up as docker container

Create environment file by copying distribution configuration:

```shell
cp .env.dist .env
```
Adjust parameters according to your environment.

Build and run image.

```shell
docker build -t shorturl .
docker run --rm -i --env-file ./.env -p8080:8080 shorturl
```

Please note that in this example container is removed as soon as it is stopped (option `--rm`). 
This example assumes than application listens on port `8080` (environment parameter `LISTEN_ADDR`).

## Frontend (in development)

`GET:/`

Public SPA where user can create shortened URL as well as get report on particular URL.  

## API
Authentication using pre-shared key provided via header `Authorization`.

```text
Authorization: bearer _pre-shared_token_
```

### Add URL

#### Request

`POST:/`

```json
{
  "url": "_origin_url_",
  "shortId": "_optional_short_id_"
}
```

#### Response

201: Created

400: Bad request, in response body JSON object with property `errorCode`

- invalid_url: Invalid URL
- invalid_short_id: Short ID already in use (only if `shortId` is non-null)


### URL report

#### Request

`GET:/view/{shortId}`

#### Response

200: Success

```json
{
  "clickCount": 0,
  "uniqueClickCount": 0
}
```

404: Short ID not found

410: Gone. Short ID was created but now disabled.

### URL redirect

#### Request

`GET:/go/{shortId}`

#### Response

Redirects to respective URL

## Environment config

- `LISTEN_ADDR` Listening address. Might be set as port only, e.g. `:8080`
- `REPOSITORY` Repository type, either `MEMORY` or `MYSQL`
  - If repository of type `MYSQL` is used:
    - `MYSQL_DBNAME` MySQL database name
    - `MYSQL_HOST` MySQL host name or IP
    - `MYSQL_USER` MySQL username
    - `MYSQL_PASS` MySQL user password
    - `MYSQL_PORT` MySQL service TCP port
- `GIN_MODE` Default is debug mode. For production environment set to `release`
- `SHARED_TOKEN` Shared token for authentication

## Tests

```shell
go test -v ./...
```

## Planned features
- ~~Redis repository~~
- Database initialization
- ~~Redis cache~~
- New URL add and redirect stats via RabbitMQ