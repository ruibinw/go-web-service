#  Go Web Service
A simple REST Web service written in Go that supports CRUD operations.

## Technology Stack
- Language:             Go 1.19
- Web Framework:        Echo
- ORM Framework:        GORM
- Dependency Injection: Wire
- Unit Testing:         Testify+GoMock
- API Documentation:    Swaggo/Swag
- Database:             MySQL
- Containerization:     Docker


## Unit Testing
To run unit tests for domains/record:
```shell
go test -v ./internal/domains/record
```
Note: Unit tests of repository layer use sqlite db that requires [gcc](https://gcc.gnu.org/install/binaries.html) to compile.

## Building and Deployment using Docker
```shell
# Build docker image
docker build -t go-web-service/api-server:latest .

# Start up services
docker-compose -f ./build/docker-compose.yml up -d
```

8080 and 3306 ports are used for the web api and database, by default.

Ports can be specified by setting environment variables `SERVER_PORT` and `DB_PORT`

## Using Makefile to simplify command lines
```shell
# run unit tests
make test

# build docker image
make build-docker

# run all docker services in docker-compose file
make run-docker
```


## API Definition
API definition of the web service can be found on url `/swagger/index.html`