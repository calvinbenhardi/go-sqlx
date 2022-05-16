# Database Operations
db:
	docker run --name gosqlx -p 5000:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it gosqlx createdb --username=root --owner=root bank

migration:
	migrate create -ext sql -dir migration -seq $(name)

migrateup:
	migrate -path migration -database postgresql://root:secret@localhost:5000/bank?sslmode=disable up

migratedown:
	migrate -path migration -database postgresql://root:secret@localhost:5000/bank?sslmode=disable down

# Development Server
dev:
	air -c .air.toml

# Test
test:
	go test -v -cover ./...