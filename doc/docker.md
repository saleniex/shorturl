Dockerized application
==

Application dockerized version can be run using docker-compose scripts:
- [With MySQL repository](../deployments/compose-mysql-repo.yaml)
- [With Redis repository](../deployments/compose-redis-repo.yaml)
- [With cached MySQL repository](../deployments/compose-cached-mysql-repo.yaml). MySQL is main data storage and Redis as caching layer.
- [Distributed repository](../deployments/compose-distributed.yaml). Same as cached MySQL with addition of access stats are stored using RabbitMQ AMQP service.

Usage example of docker compose file (from project root):

```shell
docker-compose -f ./deployments/compose-distributed.yaml up
```
