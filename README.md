# Repository for study and test microservice with Golang

## **ms-person**
### Microservice Person is golang crud rest api. This example contains metrics, healthcheck, logs, traces, swagger, database, prometheus, env

### Run Requirements
* Docker 17.6.x
* Docker Compose 1.11.x

## You can run everything in docker with one command
### This command run build ms-person image, run Api Person in docker and run Infra with one command
```
make run-all-docker
```

## Endpoints:
- Swagger: http://localhost:8080/swagger/index.html
- Jaeger: http://localhost:16686/search
- Prometheus: http://localhost:9090/graph


## For Development
### Requirements
* Golang 1.20
* Docker 17.6.x
* Docker Compose 1.11.x

### Run Only infra Containers
```
make run-infra
```

### Run DEV local Person API
```
make run-ms-person
```



