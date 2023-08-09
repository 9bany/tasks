DB_URL="postgresql://root:secret@localhost:5432/demo?sslmode=disable"
postgres:
	docker run --name postgres-latest -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres 
createdb: 
	docker exec -it postgres-latest createdb --username=root --owner=root demo
dropdb:
	docker exec -it postgres-latest dropdb --username=root demo
migrateup:
	migrate -path db/migration -database $(DB_URL) -verbose up  
migratedown:
	migrate -path db/migration -database $(DB_URL) -verbose down  
sqlc:
	sqlc generate 
mock: 
	mockgen -package mock -source db/sqlc/querier.go  -destination=./db/mock/querier.go
	mockgen -package mock -source http_client/iframely_client.go -destination=./http_client/mock/iframely.go