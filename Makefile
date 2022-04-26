run:
	go run ./cmd/main.go

.DEFAULT_GOAL := run

migrate_up:
	migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' up

migrate_down:
	migrate -path ./schema -database 'postgres://postgres:password@localhost:5436/postgres?sslmode=disable' down
