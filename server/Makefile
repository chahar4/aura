dockerinit:
	 docker run --name postgresaura -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres

createdb:
	docker exec -it postgresaura createdb --username=root --owner=root auraDB

postgres:
	docker exec -it postgresaura psql

dropdb:
	docker exec -it postgresaura dropdb auraDB 

migratecreate:
	migrate create -ext sql -dir db/migrations add_channel_table

migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/auraDB?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/auraDB?sslmode=disable" -verbose down

.PHONY: dockerinit postgres createdb dropdb migrateup migratedown

