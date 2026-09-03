include .env
export
DATABASE_URL := $(shell grep DATABASE_URL .env | cut -d '=' -f2-)
export DATABASE_URL

run-local-app:
	@ go run ./cmd/todoapp/main.go

migrate-up:
	@ ~/go/bin/migrate -path migrations -database $(DATABASE_URL) up 

migrate-down:
	@ ~/go/bin/migrate -path migrations -database $(DATABASE_URL) down

run-container-app:
	@ docker run -p 8080:8080 -d todolist_app

docker-compose-up:
	@ docker compose up -d 

docker-compose-down:
	@ docker compose down 