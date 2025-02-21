migrate:
	migrate -path ./migrations/schema -database 'postgres://postgres:root@localhost:5432/ayan?sslmode=disable' up