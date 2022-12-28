## Environment config

- `LISTEN_ADDR` Listening address. Might be set as port only, e.g. `:8080`
- `REPOSITORY` Repository type
    - If repository of type `MYSQL` is used:
        - `MYSQL_DBNAME` MySQL database name
        - `MYSQL_HOST` MySQL host name or IP
        - `MYSQL_USER` MySQL username
        - `MYSQL_PASS` MySQL user password
        - `MYSQL_PORT` MySQL service TCP port
- `GIN_MODE` Default is debug mode. For production environment set to `release`
- `SHARED_TOKEN` Shared token for authentication
- `REDIS_HOST` Redis host along with port number (localhost:6379). Relevant for repositories `REDIS`, `CACHED_MYSQL`.
- `AMQP_URL` AMQP endpoint (e.g. amqp://localhost:5762)
- `AMQP_QUEUE_NAME` AMQP queue name (default "shorturl_stats")

Available repositories:
- `MEMORY` In memory storage. Can be used for tests.
- `MYSQL` Store URLs and access stats in MySQL database. Requires `MYSQL_*` parameter configuration
- `REDIS` Store URLs in Redis database. For production use database should be configured as persistent.
- `CACHED_MYSQL` Data is stored in MySQL with Redis caching layer. Require `MYSQL_*` and `REDIS_*` parameter configuration.
- `DISTRIBUTED` Same as `CACHED_MYSQL` with addition of RabbitMQ queue for access stats storage.
