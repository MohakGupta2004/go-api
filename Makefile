DB_DOCKER_NAME=api_db
BINARY_NAME=api

postgres:
	sudo docker run --name ${DB_DOCKER_NAME} -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

run:
	go run ./cmd/server/main.go
