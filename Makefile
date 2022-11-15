unit-tests:
	go test ./...

docker-compose:
	docker compose -f ./build/docker-compose.yml build
	docker compose -f ./build/docker-compose.yml up

api-docs:
	swag i --dir cmd/server,internal/controllers,internal/models,internal/dto,internal/utils