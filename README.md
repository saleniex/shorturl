Short URL
==

URL shortening webservice with REST API for URL management. 

## Frontend

`GET:/`

Public SPA where user can create shortened URL as well as get report on particular URL. 

## API
Authenticaiton using preshared key provided via header `Authorization`.

```text
Authorization: bearer _preshared_token_
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
