VERSION=1.0.0

build-bin:
	go build -o gowebservice-v.$(VERSION) .

build-exe:
	go build -o gowebservice-v.$(VERSION).exe .

test:
	go test -v ./internal/domains/record

build-docker:
	docker build -t go-web-service/api-server:latest .

run-docker:
	docker-compose -f ./build/docker-compose.yml up

apidocs:
	swag i --dir ./,internal/domains/record,internal/models,internal/utils