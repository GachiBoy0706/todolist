include .env
export
run-local-app:
	@ go run ./cmd/todoapp/main.go

migrate-up:
	@ migrate -path migrations up