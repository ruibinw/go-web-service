#  Go Web Service
A simple REST Web service written in Go that supports CRUD operations.

## Technology Stack
- Language:             Go 1.19
- Web Framework:        Echo
- ORM Framework:        GORM
- Dependency Injection: Wire
- Unit Testing:         Testify+GoMock
- Database:             MySQL
- Containerization:     Docker

## Unit Testing
To run unit tests:
```shell
go test ./...
```


## Building and Deployment
### Build docker image
```shell
docker build -t go-web-service/api-server:latest .
```
### Start up services
```shell
docker-compose -f ./build/docker-compose.yml up -d
```

Default ports, 8080 and 3306, are used for the web api and database, respectively.

To specify ports for the web api (SERVER_PORT) and database (DB_PORT):
```shell
# Unix Shell
DB_PORT=port1 SERVER_PORT=port2 docker-compose -f ./build/docker-compose.yml up

# Windows CMD
set DB_PORT=port1
set SERVER_PORT=port2 
docker-compose -f ./build/docker-compose.yml up -d
```


## API Definition
API definition of the web service can be found on url `/swagger/index.html`