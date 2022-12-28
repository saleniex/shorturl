Short URL
==

URL shortening webservice with REST API for URL management. 

- [API](doc/api.md)
- [Environment configuration](doc/conf.md)
- [Building and testing](doc/build.md)
- [Dockerized version](doc/docker.md)

Application is implemented with CLI and can be called in two modes:
- Web service which exposes [API](doc/api.md)
- Stats consumer which consumes access stats messages from AMQP queue

## Planned features
- ~~Redis repository~~
- Database initialization
- ~~Redis cache~~
- ~~New URL add and redirect stats via RabbitMQ~~
- ~~Auto ID generation~~
