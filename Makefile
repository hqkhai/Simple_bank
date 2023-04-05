postgres:
	 docker run --name postgresforbank --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	 docker exec -it postgresforbank createdb --username=root --owner=root simple_bank

dropdb:
	 docker exec -it postgresforbank dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "C:\Users\9A312\OneDrive - VNU-HCMUS\Documents\Nam3\HKII\SimpleBank":/src -w /src kjconroy/sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -destination db/mock/store.go simplebank/db/sqlc Store   

migratedown1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrateup1:
	migrate -path db/migration -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

.PHONENY: createdb dropdb postgres migrateup migratedown sqlc test server mock migratedown1 migrateup1
