API
==

## Frontend

`GET:/`

Public SPA where user can create shortened URL as well as get report on particular URL.
(under development)

## Authentication
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
  "accessCount": 3,
  "id": "id",
  "shortIdUri": "https://go.to.address.net"
}
```

404: Short ID not found

410: Gone. Short ID was created but now disabled.

### URL redirect

#### Request

`GET:/go/{shortId}`

#### Response

Redirects to respective URL