build:
	go build .

test:
	go test -v ./internal/domains/record

build-docker:
	docker build -t go-web-service/api-server:latest .

run-docker:
	docker-compose -f ./build/docker-compose.yml up

apidocs:
	swag i --dir ./,internal/domains/record,internal/models,internal/utils