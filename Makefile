postgres:
	docker run --name postgres14 -p 5432:54322 -e POSTGRES_USER=billybills -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb -U billybills --owner=billybills billy_bills

dropdb:
	docker exec -it postgres14 dropdb -U billybills billy_bills

migrateup:
	migrate -path db/migrations -database "postgresql://billybills:secret@localhost:5432/billy_bills?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://billybills:secret@localhost:5432/billy_bills?sslmode=disable" -verbose down

sqlc:
	sqlc generate
.PHONY: postgres, createdb, dropdb, migratedown, migrateup, sqlc