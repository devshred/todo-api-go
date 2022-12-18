# Todo-API based on Go
This project uses [Go](https://go.dev/) with [Gin](https://gin-gonic.com/) and [GORM](https://gorm.io/) to provide a simple todo-app API and is used for workshops.
It implements an [OpenAPI spec](src/main/resources/todo-spec.yaml) and can be tested with a [frontend based on Vue.js](https://github.com/devshred/todo-web).

## how-to run
### start database
```shell
docker compose up
```

### run binary
#### build binary
```shell
go mod download
go build -o todo-api-app
```
#### start binary
```shell
./todo-api-app
```

### run docker container
#### build image
```shell
docker build -t todo-api-go:latest .
```
#### start container
```shell
./run.sh
```
