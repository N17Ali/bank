postgres:
	docker run --name postgres16.3 -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:16.3-alpine3.19

createdb: 
	docker exec -it postgres16.3 createdb --username=postgres --owner=postgres bank

dropdb: 
	docker exec -it postgres16.3 dropdb --username=postgres  bank

migrateup:
	migrate -path db/migration/ -database "postgres://postgres:secret@localhost:5432/bank?sslmode=disable" --verbose up

migrateup1:
	migrate -path db/migration/ -database "postgres://postgres:secret@localhost:5432/bank?sslmode=disable" --verbose up 1

migratedown:
	migrate -path db/migration/ -database "postgres://postgres:secret@localhost:5432/bank?sslmode=disable" --verbose down

migratedown1:
	migrate -path db/migration/ -database "postgres://postgres:secret@localhost:5432/bank?sslmode=disable" --verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/n17ali/bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock migratedown1 migrateup1
