Dockerized application
==

Application dockerized version can be run using docker-compose scripts:
- [With MySQL repository](../deployments/compose-mysql-repo.yaml)
- [With Redis repository](../deployments/compose-redis-repo.yaml)
- [With cached MySQL repository](../deployments/compose-cached-mysql-repo.yaml). MySQL is main data storage and Redis as caching layer.
- [Distributed repository](../deployments/compose-distributed.yaml). Same as cached MySQL with addition of access stats are stored using RabbitMQ AMQP service.

All compose files rely on defaults in file `.env`. You can use [defaults](../.env.dist) from `.env.dist`. E.g. prior to use `docker-compose`
copy distribution file and adjust according to your environment.

```shell
cp .env.dist .env
```

Usage example of docker compose file (from project root):

```shell
docker-compose -f ./deployments/compose-distributed.yaml up
```
