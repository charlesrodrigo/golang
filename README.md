# Repository for study and test microservice with Golang

# Microservice ms-person
**ms-person** is golang crud rest api. This example contains metrics, healthcheck, logs, traces, swagger, database, prometheus, env

## Development
### Requirements
* golang 1.20
* Docker 17.6.x
* Docker Compose 1.11.x

### Run infra Containers
```
make run-infra
```

### Run Person API
```
make run-ms-person
```

Access: http://localhost:8080/swagger/index.html

